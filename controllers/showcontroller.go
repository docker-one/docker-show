package controllers

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	client2 "github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DockerShow struct {
	DataMap map[string]string
}

func (d *DockerShow) GetDockerImagesList(g *gin.Context) {
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	resultData := make([]types.ImageSummary, 0)
	cli, err := client2.NewClientWithOpts(client2.WithAPIVersionNegotiation())
	if err == nil {
		resultData = listImages(cli)
	}

	rsp := new(Rsp)
	//var parms IniConfig
	//err := g.BindJSON(&parms)
	//if err != nil {
	//	rsp.Msg = "json failed"
	//	rsp.Code = 201
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
	//value := utils.GetKey("cloudserver", parms.Key)

	var reData ResultDataList
	reData.Total = int64(len(resultData))
	reData.Data = resultData

	rsp.Msg = "success"
	rsp.Code = 200
	rsp.Data = reData
	g.JSON(http.StatusOK, rsp)
}

func listImages(cli *client2.Client) []types.ImageSummary {
	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, image := range images {
		fmt.Printf("%s %s \n\n", image.ID[:19], image.RepoTags)
	}
	return images
}
func listContainers(cli *client2.Client) []types.Container {
	ty := types.ContainerListOptions{}
	ty.All = true
	cons, err := cli.ContainerList(context.Background(), ty)
	if err != nil {
		panic(err)
	}

	return cons
}

func (d *DockerShow) GetContainersList(g *gin.Context) {
	//cli, err := client.NewClientWithOpts(client.FromEnv)
	resultData := make([]types.Container, 0)
	cli, err := client2.NewClientWithOpts(client2.WithAPIVersionNegotiation())
	if err == nil {
		resultData = listContainers(cli)
	}

	rsp := new(Rsp)
	//var parms IniConfig
	//err := g.BindJSON(&parms)
	//if err != nil {
	//	rsp.Msg = "json failed"
	//	rsp.Code = 201
	//	g.JSON(http.StatusOK, rsp)
	//	return
	//}
	//value := utils.GetKey("cloudserver", parms.Key)

	var reData ResultDataList
	reData.Total = int64(len(resultData))
	reData.Data = resultData

	rsp.Msg = "success"
	rsp.Code = 200
	rsp.Data = reData
	g.JSON(http.StatusOK, rsp)
}
