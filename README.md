## Text2Voice 

Text2Voice是一个将文本信息转成语音的小工具，基于讯飞语音SDK实现的。
因为讯飞只提供了Linux C的SDK，所以此工具的运行环境是在Linux下。


### 环境依赖

* go version >= 1.11, go mod


### 使用方式

```bash 
git clone git@github.com:dyike/text2voice.git
cd $GOPATH/src/text2voice 
go mod tidy

# 将libmsc.so加入环境变量
# 64位动态库
mv xf/libs/x64/libmsc.so /usr/local/lib/
# 32位动态库
mv xf/libs/x86/libmsc.so /usr/local/lib/

export LD_LIBRARY_PATH=/usr/local/lib

# 编译
go build 

# 配置文件
cp default.json config.json
```

### 配置文件

```json
{
    "log":{
        "enable":false,
        "filename": "voice.log",
        "level": "ll"
    },

    "appid": "5808ae7e",
    "work_dir": ".",
    "speed": 1,
    "tts": {
        "voice_name":"aisjiuxu", 
        "text_encoding":"utf8", 
        "sample_rate":16000,
        "speed":50, 
        "volume":50, 
        "pitch":50, 
        "rdn":2
    }
}
```
* log.enable 是否开启日志
* log.filename 日志文件名称
* log.level 日志等级
* appid 讯飞开放平台注册申请的应用的appid,请换成自己的appid才能正常使用
* work_dir 工作目录，默认当前目录下
* speed 合成速度，默认1，范围1-10，数值越小速度越快越耗CPU
* tts.voice_name 发音人名称，默认aisjiuxu(讯飞许久),支持:xiaoyan,aisxping,aisjinger,aisbabyxu,等其他付费
* tts.text_encoding 文本编码格式，默认为utf8
* tts.sample_rate 音频采样率，默认16000 
* tts.speed 发音人语速，默认50
* tts.volume 发音人音量，默认50
* tts.pitch 发音人音调， 默认50 
* tts.rdn 发音方式


### 使用示例

```
# 合成文件内容
./text2voice -c config.json -p test.txt -o test.mp3

# 合成文本
./text2voice -c config.json -t "哈哈，就是这样使用嘀~~" -o test.mp3
```

### Enjoy~~~
