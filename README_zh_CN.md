# CheckNatType_GUI

CheckNatType_GUI 是一个使用 Go 语言和 Fyne 框架开发的图形用户界面 (GUI) 应用程序，用于检测用户的 NAT 类型。该程序可以选择内置的 STUN 服务器或自行输入 STUN 服务器地址进行检测。

## 特性

- 使用 Go 语言编写，性能高效
- 友好的图形用户界面，易于使用
- 支持多种内置 STUN 服务器
- 允许用户输入自定义的 STUN 服务器地址
- 支持多种操作系统和架构

## 内置的 STUN 服务器

程序内置了以下 STUN 服务器供用户选择：

- 国际 STUN 服务器
  - `stun.l.google.com:19302`
  - `stun1.l.google.com:19302`
  - `stun2.l.google.com:19302`
  - `stun3.l.google.com:19302`
  - `stun4.l.google.com:19302`
  - `stun.ekiga.net:3478`
  - `stun.ideasip.com:3478`
  - `stun.schlund.de:3478`
  - `stun.stunprotocol.org:3478`
  - `stun.voipbuster.com:3478`
  - `stun.voipstunt.com:3478`
  - `stun.services.mozilla.com:3478`

- 中国大陆 STUN 服务器
  - `stun.qq.com:3478`
  - `stun.miwifi.com:3478`
  - `stun.wide.net.cn:3478`
  - `stun.chat.bilibili.com:3478`

- 其他 STUN 服务器
  - `turn.cloudflare.com:3478`
  - `fwa.lifesizecloud.com:3478`
  - `stun.isp.net.au:3478`
  - `stun.freeswitch.org:3478`
  - `stun.voip.blackberry.com:3478`
  - `stun.nextcloud.com:3478`
  - `stun.sipnet.com:3478`
  - `stun.radiojar.com:3478`
  - `stun.sonetel.com:3478`
  - `stun.voipgate.com:3478`
  - `stun.cloudflare.com:53`
  - `stun.cloudflare.com:80`

## 下载与安装

你可以从 GitHub 的 [Releases 页面](https://github.com/YOUR_USERNAME/CheckNatType_GUI/releases) 下载对应操作系统和架构的二进制文件：

- Windows 64位: `CheckNatType_GUI_windows_amd64.exe`
- Windows 32位: `CheckNatType_GUI_windows_386.exe`
- Linux 64位: `CheckNatType_GUI_linux_amd64`
- Linux 32位: `CheckNatType_GUI_linux_386`
- Linux ARM 64位: `CheckNatType_GUI_linux_arm64`
- Linux ARM 32位: `CheckNatType_GUI_linux_arm`
- macOS 64位: `CheckNatType_GUI_darwin_amd64`
- macOS ARM 64位: `CheckNatType_GUI_darwin_arm64`

### 在 Windows 上运行

1. 下载 `CheckNatType_GUI_windows_amd64.exe` 或 `CheckNatType_GUI_windows_386.exe`。
2. 双击运行下载的文件。

### 在 Linux 上运行

1. 下载适用于你的架构的二进制文件，例如 `CheckNatType_GUI_linux_amd64`。
2. 为文件添加可执行权限：

    ```bash
    chmod +x CheckNatType_GUI_linux_amd64
    ```

3. 运行程序：

    ```bash
    ./CheckNatType_GUI_linux_amd64
    ```

### 在 macOS 上运行

1. 下载适用于你的架构的二进制文件，例如 `CheckNatType_GUI_darwin_amd64`。
2. 为文件添加可执行权限：

    ```bash
    chmod +x CheckNatType_GUI_darwin_amd64
    ```

3. 运行程序：

    ```bash
    ./CheckNatType_GUI_darwin_amd64
    ```

## 使用方法

1. 打开应用程序后，你可以选择一个内置的 STUN 服务器，或输入自定义的 STUN 服务器地址。
2. 点击 "Check NAT Type" 按钮进行检测。
3. 检测结果会显示在界面上。

## 编译

确保你已经安装了 Go 语言环境和 Git 工具。然后执行以下步骤：

1. 克隆仓库：

    ```bash
    git clone https://github.com/YOUR_USERNAME/CheckNatType_GUI.git
    cd CheckNatType_GUI
    ```

2. 编译项目：

    在 Windows 上：

    ```bash
    go build -ldflags "-H=windowsgui" -o CheckNatType_GUI.exe main.go
    ```

    在 Linux 和 macOS 上：

    ```bash
    go build -o CheckNatType_GUI main.go
    ```

## 贡献

欢迎贡献代码！

## 许可证

本项目使用 GPL-3.0  许可证，详情请参阅 `LICENSE` 文件。

## 联系

如果你有任何问题或建议，请联系项目维护者：

- 邮箱: support@zyh8.com
- My Blog: [https://www.zyh8.com](https://www.zyh8.com)

## 捐赠
如果您喜欢这个项目，可以捐赠到以下地址：
USDT (TRC20): TJsH3fGwtmr1nSbyBfFN6uXvBArpQjebJ6

## 图
![image](https://github.com/user-attachments/assets/b23a0501-4d8e-43ff-bf02-6a6ff669e78d)
![image](https://github.com/user-attachments/assets/57565d76-030f-4499-9725-581c61fa8029)
![image](https://github.com/user-attachments/assets/42737ee4-f25e-4302-ad4e-b92b87154301)
