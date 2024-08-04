# CheckNatType_GUI

CheckNatType_GUI is a graphical user interface (GUI) application developed in Go using the Fyne framework to detect the user's NAT type. The program allows you to select from built-in STUN servers or input custom STUN server addresses for detection.

**Languages: [English](#english) | [简体中文](https://github.com/zhiyunhai/CheckNatType_GUI/blob/main/README_zh_CN.md)**

---

## English

## Features

- Developed in Go for high performance
- User-friendly graphical interface
- Supports multiple built-in STUN servers
- Allows input of custom STUN server addresses
- Supports multiple operating systems and architectures

## Built-in STUN Servers

The program includes the following STUN servers for users to choose from:

- International STUN Servers
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

- China STUN Servers
  - `stun.qq.com:3478`
  - `stun.miwifi.com:3478`
  - `stun.wide.net.cn:3478`
  - `stun.chat.bilibili.com:3478`

- Other STUN Servers
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

## Download and Install

You can download the binary files for your operating system and architecture from the [Releases page](https://github.com/YOUR_USERNAME/CheckNatType_GUI/releases) on GitHub:

- Windows 64-bit: `CheckNatType_GUI_windows_amd64.exe`
- Windows 32-bit: `CheckNatType_GUI_windows_386.exe`
- Linux 64-bit: `CheckNatType_GUI_linux_amd64`
- Linux 32-bit: `CheckNatType_GUI_linux_386`
- Linux ARM 64-bit: `CheckNatType_GUI_linux_arm64`
- Linux ARM 32-bit: `CheckNatType_GUI_linux_arm`
- macOS 64-bit: `CheckNatType_GUI_darwin_amd64`
- macOS ARM 64-bit: `CheckNatType_GUI_darwin_arm64`

### Running on Windows

1. Download `CheckNatType_GUI_windows_amd64.exe` or `CheckNatType_GUI_windows_386.exe`.
2. Double-click the downloaded file to run it.

### Running on Linux

1. Download the binary file for your architecture, e.g., `CheckNatType_GUI_linux_amd64`.
2. Add execute permission to the file:

    ```bash
    chmod +x CheckNatType_GUI_linux_amd64
    ```

3. Run the program:

    ```bash
    ./CheckNatType_GUI_linux_amd64
    ```

### Running on macOS

1. Download the binary file for your architecture, e.g., `CheckNatType_GUI_darwin_amd64`.
2. Add execute permission to the file:

    ```bash
    chmod +x CheckNatType_GUI_darwin_amd64
    ```

3. Run the program:

    ```bash
    ./CheckNatType_GUI_darwin_amd64
    ```

## Usage

1. Open the application and select a built-in STUN server or input a custom STUN server address.
2. Click the "Check NAT Type" button to start detection.
3. The detection result will be displayed on the interface.

## Building

Ensure you have Go and Git installed. Then follow these steps:

1. Clone the repository:

    ```bash
    git clone https://github.com/YOUR_USERNAME/CheckNatType_GUI.git
    cd CheckNatType_GUI
    ```

2. Build the project:

    On Windows:

    ```bash
    go build -ldflags "-H=windowsgui" -o CheckNatType_GUI.exe main.go
    ```

    On Linux and macOS:

    ```bash
    go build -o CheckNatType_GUI main.go
    ```

## Contributing

Contributions are welcome!

## License

This project is licensed under the GPL-3.0 License. See the `LICENSE` file for details.

## Contact

If you have any questions or suggestions, please contact the project maintainer:

- Email: support@zyh8.com
- My Blog: [https://www.zyh8.com](https://www.zyh8.com)

## Donations
If you like this project, you can donate to the following address:
USDT (TRC20): TJsH3fGwtmr1nSbyBfFN6uXvBArpQjebJ6

## Screenshots
![image](https://github.com/user-attachments/assets/b23a0501-4d8e-43ff-bf02-6a6ff669e78d)
![image](https://github.com/user-attachments/assets/57565d76-030f-4499-9725-581c61fa8029)
![image](https://github.com/user-attachments/assets/42737ee4-f25e-4302-ad4e-b92b87154301)
