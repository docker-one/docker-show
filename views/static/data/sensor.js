function methodInfo(methodName,data){
    this.method=methodName
    this.data=data
}
let serverOrigin=location.origin;
let serverUrl=serverOrigin+'/v1';
//let serverUrl='http://localhost:9001/v1'
let code_success=200;
var sensorServer= {
    //向服务器法送数据
    postServerData: function (postUrl, postData) {
        $.ajaxSettings.async = false;
        //this.filterMineByInfo(postData);
        var url = postUrl;
        var serverResultData;
        var jsonData = JSON.stringify(postData)
        $.ajax({
            //请求类型，这里为POST
            // xhrFields:{
            //     withCredentials:true
            // },
            type: 'POST', //你要请求的api的URL
            url: url, //是否使用缓存
            cache: false, //数据类型，这里我用的是json
            dataType: "json",
            contentType: "application/json;charset=utf-8",
            //必要的时候需要用JSON.stringify() 将JSON对象转换成字符串
            data: jsonData, //data: {key:value},
            //添加额外的请求头
            headers: {'Access-Control-Allow-Origin': '*', "token": window.localStorage.getItem("token")}, //请求成功的回调函数
            success: function (data) {
                //函数参数 "data" 为请求成功服务端返回的数据
                // if (data.code==500&&data.msg==401){
                //     //console.log("重新登录！");
                //     layer.msg("请重新登陆！",{icon:1,time:1000});
                //     //window.open("./login.html");
                //     //window.open("./login.html", "_blank", "resizable,scrollbars,status");
                //     top.location.href ="login.html"; //这种可以新页打开
                //     //window.parent.location.replace("login.html");//刷新父级页面; 这种可以新页打开
                // }
                serverResultData = data.data;
            }
        });
        $.ajaxSettings.async = true;
        return serverResultData;
    },
    postServerDataSync: function (postUrl, postData) {
        //   $.ajaxSettings.async = false;
        var url = postUrl;
        var serverResultData = false;
        var jsonData = JSON.stringify(postData)
        $.ajax({
            //请求类型，这里为POST
            type: 'POST', //你要请求的api的URL
            url: url, //是否使用缓存
            cache: false, //数据类型，这里我用的是json
            dataType: "json",
            contentType: "application/json;charset=utf-8",
            //必要的时候需要用JSON.stringify() 将JSON对象转换成字符串
            data: jsonData, //data: {key:value},
            //添加额外的请求头
            headers: {'Access-Control-Allow-Origin': '*', "token": window.localStorage.getItem("token")}, //请求成功的回调函数
            success: function (data) {
                serverResultData = data.data;
                return data.data;
            }
        });
        // $.ajaxSettings.async = true;
        return serverResultData;
    },
    //获得已经登录人的信息
    getLoginUserInfo:function(){
        //return localStorage.getItem("robotUser");
        return JSON.parse(localStorage.getItem("robotUser"));
    },
    //加过滤矿 使用 this.filterMineByInfo(postData);
    filterMineByInfo:function(postData){
        //如果有就别管
        if(postData.hasOwnProperty("minesId"))
            return;
        if(this.getLoginUserInfo()==null)
            postData.minesId=0;
        else
            postData.minesId=this.getLoginUserInfo().minesId;
        //return postData;
    },

    getImagesList: function (postData) {
        //debugger;
        //this.filterMineByInfo(postData);
        var url = serverUrl + "/docker/imagesList";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },

    getContainersList: function (postData) {
        //debugger;
        //this.filterMineByInfo(postData);
        var url = serverUrl + "/docker/containersList";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    //修改传感器数据
    updateSensorInfo1: function (postData) {
        $.ajaxSettings.async = false;
        var url = serverUrl + "/sensor/update";
        return sensorServer.postServerData(url, postData);
    },
    //新增传感器数据
    addSensorInfo: function (postData) {
        // var url=serUrl;
        // var postData=new methodInfo("AddSensor",jsonData)
        // jsonData=JSON.stringify(postData)
        // $.post(url,jsonData,function (data) {
        //     console.log(data);
        //     return data;
        // },'json')
        // return null;
        $.ajaxSettings.async = false;
        var url = serverUrl + "/sensor/addSensor";
        return sensorServer.postServerData(url, postData);
    },
    delSensorInfo: function (postData) {
        $.ajaxSettings.async = false;
        var url = serverUrl + "/sensor/delete";
        return sensorServer.postServerData(url, postData);
    },

    //-----------------------
    //获取传感器数据 list
    getRealTimeGas: function (postData) {
        //debugger;
        //this.filterMineByInfo(postData);
        var url = serverUrl + "/led/getRealTimeGas";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    getCurrentData: function (postData) {
        var url = serverUrl + "/led/getCurrentData";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    saveCurrentData: function (postData) {
        var url = serverUrl + "/led/saveCurrentData";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    getConfigInfo: function (postData) {
        var url = serverUrl + "/led/getConfigInfo";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
    saveConfigInfo: function (postData) {
        var url = serverUrl + "/led/saveConfigInfo";
        $.ajaxSettings.async = false;
        return sensorServer.postServerData(url, postData);
    },
}

function downLoadByUrl(url){
    var xhr = new XMLHttpRequest();
    //GET请求,请求路径url,async(是否异步)
    xhr.open('GET', url, true);
    //设置请求头参数的方式,如果没有可忽略此行代码
    xhr.setRequestHeader("token", window.localStorage.getItem("token"));
    //设置响应类型为 blob
    xhr.responseType = 'blob';
    //关键部分
    xhr.onload = function (e) {
        //如果请求执行成功
        if (this.status == 200) {
            var blob = this.response;
            var filename = "";//如123.xls
            var a = document.createElement('a');

            blob.type = "application/octet-stream";
            //创键临时url对象
            var url = URL.createObjectURL(blob);

            a.href = url;
            a.download=filename;
            a.click();
            //释放之前创建的URL对象
            window.URL.revokeObjectURL(url);
        }
    };
    //发送请求
    xhr.send();
}


