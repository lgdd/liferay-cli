package scaffold

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/iancoleman/strcase"

	"github.com/lgdd/lfr-cli/pkg/metadata"
	"github.com/lgdd/lfr-cli/pkg/util/fileutil"
	"github.com/lgdd/lfr-cli/pkg/util/logger"
)

// ServiceBuilderData contains the data to be injected into the template files
type ServiceBuilderData struct {
	Package                 string
	Name                    string
	CamelCaseName           string
	WorkspaceName           string
	WorkspaceCamelCaseName  string
	WorkspacePackage        string
	MajorVersion            string
	DtdMajorVersion         string
	WorkspaceProductEdition string
}

// Creates the structure for a Service Builder module
func CreateModuleServiceBuilder(liferayWorkspace, name string) {
	sep := string(os.PathSeparator)

	modulePackage := metadata.PackageName
	workspacePackage, _ := metadata.GetGroupId()

	if modulePackage == "org.acme" && workspacePackage != "org.acme" {
		modulePackage = strings.Join([]string{workspacePackage, strcase.ToDelimited(name, '.')}, ".")
	}

	destModuleParentPath := filepath.Join(liferayWorkspace, "modules")
	destModulePath := filepath.Join(destModuleParentPath, name)
	destModuleAPIPath := filepath.Join(destModuleParentPath, name, name+"-api")
	destModuleServicePath := filepath.Join(destModuleParentPath, name, name+"-service")
	camelCaseName := strcase.ToCamel(name)
	workspaceSplit := strings.Split(liferayWorkspace, sep)
	workspaceName := workspaceSplit[len(workspaceSplit)-1]

	err := fileutil.CreateDirsFromAssets("tpl/service_builder", destModulePath)

	if err != nil {
		logger.Fatal(err.Error())
	}

	err = fileutil.CreateFilesFromAssets("tpl/service_builder", destModulePath)

	if err != nil {
		logger.Fatal(err.Error())
	}

	err = renameModuleServiceBuilderFiles(destModulePath, destModuleAPIPath, destModuleServicePath)

	if err != nil {
		logger.Fatal(err.Error())
	}

	if fileutil.IsGradleWorkspace(liferayWorkspace) {
		pomPath := filepath.Join(destModulePath, "pom.xml")
		err = os.Remove(pomPath)

		if err != nil {
			logger.Fatal(err.Error())
		}

		pomPath = filepath.Join(destModuleAPIPath, "pom.xml")
		err = os.Remove(pomPath)

		if err != nil {
			logger.Fatal(err.Error())
		}

		pomPath = filepath.Join(destModuleServicePath, "pom.xml")
		err = os.Remove(pomPath)

		if err != nil {
			logger.Fatal(err.Error())
		}
	}

	if fileutil.IsMavenWorkspace(liferayWorkspace) {
		buildGradlePath := filepath.Join(destModuleAPIPath, "build.gradle")
		err = os.Remove(buildGradlePath)

		if err != nil {
			logger.Fatal(err.Error())
		}

		buildGradlePath = filepath.Join(destModuleServicePath, "build.gradle")
		err = os.Remove(buildGradlePath)

		if err != nil {
			logger.Fatal(err.Error())
		}

		pomParentPath := filepath.Join(destModulePath, "../pom.xml")
		pomParent, err := os.Open(pomParentPath)
		if err != nil {
			logger.Fatal(err.Error())
		}
		defer pomParent.Close()

		byteValue, _ := io.ReadAll(pomParent)

		var pom fileutil.Pom
		err = xml.Unmarshal(byteValue, &pom)

		if err != nil {
			logger.Fatal(err.Error())
		}

		modules := append(pom.Modules.Module, name)
		pom.Modules.Module = modules
		pom.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
		pom.SchemaLocation = "http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd"

		finalPomBytes, _ := xml.MarshalIndent(pom, "", "  ")

		err = os.WriteFile(pomParentPath, []byte(fileutil.XMLHeader+string(finalPomBytes)), 0644)

		if err != nil {
			logger.Fatal(err.Error())
		}

		logger.PrintWarn("modified ")
		fmt.Printf("%s\n", pomParentPath)
	}

	productVersion, err := fileutil.GetLiferayWorkspaceProductVersion(liferayWorkspace)

	if err != nil {
		logger.Fatal(err.Error())
	}

	var majorVersionBuilder strings.Builder
	majorVersionBuilder.WriteString(productVersion)
	majorVersionBuilder.WriteString(".0")

	workspaceProductEdition, err := fileutil.GetLiferayWorkspaceProductEdition(liferayWorkspace)

	if err != nil {
		logger.Fatal(err.Error())
	}

	data := &ServiceBuilderData{
		Package:                 modulePackage,
		Name:                    name,
		CamelCaseName:           camelCaseName,
		WorkspaceName:           workspaceName,
		WorkspaceCamelCaseName:  strcase.ToCamel(workspaceName),
		WorkspacePackage:        workspacePackage,
		MajorVersion:            majorVersionBuilder.String(),
		DtdMajorVersion:         strings.ReplaceAll(majorVersionBuilder.String(), ".", "_"),
		WorkspaceProductEdition: workspaceProductEdition,
	}

	err = updateModuleServiceBuilderWithData(destModulePath, data)

	if err != nil {
		logger.Fatal(err.Error())
	}

	_ = filepath.Walk(destModulePath,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				logger.PrintSuccess("created ")
				fmt.Printf("%s\n", path)
			}
			return nil
		})
}

func renameModuleServiceBuilderFiles(destModulePath string, destModuleAPIPath string, destModuleServicePath string) error {
	err := os.Rename(filepath.Join(destModulePath, "api"), destModuleAPIPath)

	if err != nil {
		return err
	}

	err = os.Rename(filepath.Join(destModulePath, "service"), destModuleServicePath)

	if err != nil {
		return err
	}

	err = os.Rename(filepath.Join(destModuleAPIPath, "gitignore"), filepath.Join(destModuleAPIPath, ".gitignore"))

	if err != nil {
		return err
	}

	err = os.Rename(filepath.Join(destModuleServicePath, "gitignore"), filepath.Join(destModuleServicePath, ".gitignore"))

	if err != nil {
		return err
	}

	return err
}

func updateModuleServiceBuilderWithData(destModulePath string, data *ServiceBuilderData) error {
	return filepath.Walk(destModulePath, func(path string, info fs.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !info.IsDir() {
			err = fileutil.UpdateWithData(path, data)
		}

		if err != nil {
			return err
		}

		return nil
	})
}
