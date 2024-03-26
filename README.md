# go-musicfox

go-musicfox是用Go写的又一款网易云音乐命令行客户端，支持UnblockNeteaseMusic、各种音质级别、lastfm、MPRIS、MacOS交互响应（睡眠暂停、蓝牙耳机连接断开响应、菜单栏控制等）...

> UI基于 [charmbracelet/bubbletea](https://github.com/charmbracelet/bubbletea) ，做了一些定制

![GitHub repo size](https://img.shields.io/github/repo-size/go-musicfox/go-musicfox) ![GitHub](https://img.shields.io/github/license/go-musicfox/go-musicfox) ![Last Tag](https://badgen.net/github/tag/go-musicfox/go-musicfox) ![GitHub last commit](https://badgen.net/github/last-commit/go-musicfox/go-musicfox) ![GitHub All Releases](https://img.shields.io/github/downloads/go-musicfox/go-musicfox/total)

![GitHub stars](https://img.shields.io/github/stars/go-musicfox/go-musicfox?style=social) ![GitHub forks](https://img.shields.io/github/forks/go-musicfox/go-musicfox?style=social)

<p><img src="previews/logo.png" alt="logo" width="512"/></p>

([The icon](https://github.com/go-musicfox/go-musicfox-icon) is based on [kitty-icon](https://github.com/DinkDonk/kitty-icon))

## 预览

#### 1. 启动

![启动界面](previews/boot.png)

#### 2. 主界面

![主界面](previews/main.png)

#### 3. 通知

![通知](previews/notify.png)

#### 4. 登录

![登录界面](previews/login.png)

#### 5. 搜索

![搜索界面](previews/search.png)

#### 6. lastfm授权

![lastfm](previews/lastfm.png)

#### 7. Mac NowPlaying

![NowPlaying](previews/nowplaying.png)

#### 8. UnblockNeteaseMusic

![UNM](previews/unm.png)

#### 9. MacOS歌词显示

![LyricsX](previews/lyricsX.gif)

> 1. 需要下载安装[LyricsX go-musicfox Fork版本](https://github.com/go-musicfox/LyricsX/releases/latest)
> 2. 且go-musicfox >= v3.7.7
> 3. 在LyricsX设置中，打开`使用系统正在播放的应用`

## 安装

**请务必使用等宽字体，或将配置doubleColumn设为false，否则双列显示很乱**


**Mac原生Terminal及Windows的cmd不做兼容处理**

> Mac推荐使用Iterm2或Kitty 
> 
> Linux推荐Kitty
> 
> Windows推荐使用Windows Terminal，UI及体验好很多

### Mac

#### 1. 使用Homebrew安装

```sh
brew install anhoder/go-musicfox/go-musicfox
```

如果你之前安装过musicfox，需要使用下列命令重新链接:

```sh
brew unlink musicfox && brew link --overwrite go-musicfox
```

#### 2. 直接下载

下载Mac可执行文件: https://github.com/go-musicfox/go-musicfox/releases/latest

### Linux

#### 1. 使用Homebrew安装

```sh
brew install anhoder/go-musicfox/go-musicfox
```

如果你之前安装过musicfox，需要使用下列命令重新链接:

```sh
brew unlink musicfox && brew link --overwrite go-musicfox
```

#### 2. Arch Linux可使用AUR安装：

```sh
#编译
paru -S go-musicfox
#二进制包
paru -S go-musicfox-bin
```

#### 3. Gentoo Linux可使用gentoo-zh overlay安装：

```
eselect repository enable gentoo-zh
emerge --sync
emerge -a media-sound/go-musicfox
```

#### 4. NixOS可从[Nixpkgs](https://search.nixos.org/packages?channel=unstable&show=go-musicfox&from=0&size=50&sort=relevance&type=packages&query=go-musicfox)安装：

安装到本地profile：
```
nix-env -iA nixos.go-musicfox
```

临时安装：
```
nix-shell -p go-musicfox
```

安装到Configuration.nix（推荐）：
```nix
# configuration
environment.systemPackages = [
  pkgs.go-musicfox
];

# 或者home manager
home.packages = [
  pkgs.go-musicfox
];
```

#### 5. Void Linux可使用 void-packages-zh 安装：

具体安装请看[这里](https://github.com/voidlinux-zh-association/void-packages-zh#readme)。

#### 6. 直接下载

下载Linux可执行文件: https://github.com/go-musicfox/go-musicfox/releases/latest

### Windows

**如果在使用时出现莫名奇妙的光标移动、切歌、暂停等现象，请把配置项enableMouseEvent设置为false**

#### 1. scoop安装

```sh
scoop bucket add go-musicfox https://github.com/go-musicfox/go-musicfox.git

scoop install go-musicfox
```

#### 2. 直接下载

下载Windows可执行文件: https://github.com/go-musicfox/go-musicfox/releases/latest

### 手动编译

```sh
> git clone https://github.com/go-musicfox/go-musicfox

> go mod download

(Linux编译前需要安装flac)

> make # 编译到bin目录下
> make install # 安装到$GOPATH/bin下
```

## 使用

```sh
$ musicfox
```

|    按键     |       作用       |             备注              |
|:---------:|:--------------:|:---------------------------:|
| h/H/LEFT  |       左        |                             |
| l/L/RIGHT |       右        |                             |
|  k/K/UP   |       上        |                             |
| j/J/DOWN  |       下        |                             |
|     g     |     上移到顶部      |                             |
|     G     |     下移到底部      |                             |
|    q/Q    |       退出       |                             |
|   space   |     暂停/播放      |                             |
|     [     |      上一曲       |                             |
|     ]     |      下一曲       |                             |
|   -/滚轮下   |      减小音量      |                             |
|   =/滚轮上   |      加大音量      |                             |
| n/N/ENTER |    进入选中的菜单     |                             |
|  b/B/ESC  |     返回上级菜单     |                             |
|    w/W    |    退出并退出登录     |                             |
|     p     |     切换播放方式     |                             |
|     P     | 心动模式(仅在歌单中时有效) |                             |
|    r/R    |     重新渲染UI     | 如果UI界面因为某种原因出现错乱，可以使用这个重新渲染 |
|    c/C    |     当前播放列表     |                             |
|    v/V    |    快进5s/10s    |                             |
|    x/X    |    快退1s/5s     |                             |
|     ,     |    喜欢当前播放歌曲    |                             |
|     <     |    喜欢当前选中歌曲    |                             |
|     .     |  当前播放歌曲移除出喜欢   |                             |
|     >     |  当前选中歌曲移除出喜欢   |                             |
|     t     |  标记当前播放歌曲为不喜欢  |                             |
|     T     |  标记当前选中歌曲为不喜欢  |                             |
|     d     |    下载当前播放歌曲    |                             |
|     D     |    下载当前选中歌曲    |                             |
|     /     |     搜索当前列表     |                             |
|     ?     |      帮助信息      |                             |
|     a     |   播放中歌曲的所属专辑   |                             |
|     A     |   选中歌曲的所属专辑    |                             |
|     s     |   播放中歌曲的所属歌手   |                             |
|     S     |   选中歌曲的所属歌手    |                             |
|     o     |   网页打开播放中歌曲    |                             |
|     O     | 网页打开选中歌曲/专辑... |                             |
|    ;/:    |     收藏选中歌单     |                             |
|    '/"    |    取消收藏选中歌单    |                             |

## 配置文件

配置文件路径为用户目录下的.go-musicfox/go-musicfox.ini，相关配置：

[详细配置](./utils/embed/go-musicfox.ini)

## 相关项目

1. [anhoder/bubbletea](https://github.com/anhoder/bubbletea): 基于 [bubbletea](https://github.com/charmbracelet/bubbletea) 进行部分定制 
2. [anhoder/bubbles](https://github.com/anhoder/bubbles): 基于 [bubbles](https://github.com/charmbracelet/bubbles) 进行部分定制
3. [anhoder/netease-music](https://github.com/anhoder/netease-music): fork自 [NeteaseCloudMusicApiWithGo](https://github.com/sirodeneko/NeteaseCloudMusicApiWithGo) ，在原项目的基础上去除API功能，只保留service、util作为一个独立的包，方便在其他go项目中调用

## 感谢

感谢以下项目及其贡献者们（但不限于）：

* [bubbletea](https://github.com/charmbracelet/bubbletea)
* [beep](https://github.com/faiface/beep)
* [musicbox](https://github.com/darknessomi/musicbox)
* [NeteaseCloudMusicApi](https://github.com/Binaryify/NeteaseCloudMusicApi)
* [NeteaseCloudMusicApiWithGo](https://github.com/sirodeneko/NeteaseCloudMusicApiWithGo)
* [gcli](https://github.com/gookit/gcli)
* ...

感谢[JetBrains Open Source](https://www.jetbrains.com/zh-cn/opensource/?from=archery)为项目提供免费的 IDE 授权    
[<img src="https://resources.jetbrains.com/storage/products/company/brand/logos/jb_beam.png" width="200"/>](https://www.jetbrains.com/opensource/)

[![Star History Chart](https://api.star-history.com/svg?repos=go-musicfox/go-musicfox&type=Date)](https://star-history.com/#go-musicfox/go-musicfox&Date)


补充mobian的依赖：

Installing Linux dependencies

This page lists the required dependencies to build a Bevy project on your Linux machine.

If you don't see your distro present in the list, feel free to add the instructions in this document.
Ubuntu

sudo apt-get install g++ pkg-config libx11-dev libasound2-dev libudev-dev libxkbcommon-x11-0

if using Wayland, you will also need to install

sudo apt-get install libwayland-dev libxkbcommon-dev

Depending on your graphics card, you may have to install one of the following: vulkan-radeon, vulkan-intel, or mesa-vulkan-drivers

Compiling with clang is also possible - replace the g++ package with clang.

go build github.com/cocoonlife/goflac:

# pkg-config --cflags  -- flac

Package flac was not found in the pkg-config search path.

Perhaps you should add the directory containing `flac.pc'

to the PKG_CONFIG_PATH environment variable

Package 'flac', required by 'virtual:world', not found

pkg-config: exit status 1

go build github.com/hajimehoshi/oto:

# pkg-config --cflags  -- alsa

Package alsa was not found in the pkg-config search path.

Perhaps you should add the directory containing `alsa.pc'

to the PKG_CONFIG_PATH environment variable

Package 'alsa', required by 'virtual:world', not found

pkg-config: exit status 1


sudo apt-get -y install libflac-dev flac libasound2-dev 




Backup your ~/.asoundrc - if you have one - and add the following, to a new ~/.asoundrc:

pcm.!default { 
 type plug 
 slave { 
 pcm "hw:0,0" 
 } 
} 

在 Alpine Linux arm64 上使用 go-musicfox 且经常出现 panic 的情况，可能是因为缺少了某些依赖库。要解决这个问题，可以尝试以下步骤：

    确保您的 Alpine Linux arm64 系统已经安装了 libasound2-dev 或 libasound2 库，这是 Beep 所需要的 ALSA 库。您可以使用以下命令安装：

apk add alsa-lib alsa-utils

    确保您的系统已经安装了 libasound-dev 库，这是 Beep 所需要的 ALSA 开发包。您可以使用以下命令安装：

apk add alsa-lib-dev

    确保您的系统已经安装了 libasound2-plugins 库，这是 Beep 所需要的 ALSA 插件。您可以使用以下命令安装：

apk add alsa-plugins

    如果您使用的是 Go 编程语言，还需要确保您的 Go 项目中已经正确导入了 Beep 库。您可以使用以下命令安装 Beep：

go get github.com/faiface/beep

通过以上步骤，您应该能够在 Alpine Linux arm64 上成功运行 go-musicfox，并避免出现 panic 的情况。如果问题仍然存在，请提供更多详细信息，以便我能够更好地帮助您解决问题。
