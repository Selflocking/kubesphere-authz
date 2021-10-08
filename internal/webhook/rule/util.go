package rule

import (
	"fmt"
	"ksauth/internal/config"
	"ksauth/pkg/crdadaptor"
	"strings"
)


func getAdaptorObject(policy string)(interface{},error){
	scheme,path,err:=splitSchemeAndPath(policy)
	if err!=nil{
		return nil,err
	}
	switch scheme{
	case "file":
		return path,nil
	case "crd":
		tmp:=strings.Split(path,"#")
		if len(tmp)==0{
			return nil,fmt.Errorf("invalid syntax for crd url path. correct syntax: <yaml path to crd definition>#<namespace>")
		}
		yamlPath:=tmp[0]
		namespace:=tmp[1]
		adaptor, err := crdadaptor.NewK8sCRDAdaptorByYamlDefinition(namespace, yamlPath, config.CLIENT_MODE)
		if err!=nil{
			return nil,err
		}
		return adaptor,nil
	}
	return nil,fmt.Errorf("invalid scheme %s",scheme)
}


func splitSchemeAndPath(url string)(scheme, path string, e error){
	tmp:=strings.Split(url,"://")
	scheme=""
	path=""
	if len(tmp)!=2{
		e=fmt.Errorf("invalid url %s",url)
		return 
	}
	return tmp[0],tmp[1],nil
}