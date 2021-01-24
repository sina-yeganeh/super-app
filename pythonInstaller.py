import os

try:
    import platform
except Exception as err:
    os.system("pip install platform")
    import platform

def check_system():
    system_platform = platform.system()

    if system_platform == "windows":
        return 1
    elif system_platform == "linux":
        return 0
    else:
        return 2

def install_go_on_linux():
    try:
        os.system("pkg install go")
    except Exception as err:
        print("There is a problem", err, "\nPlaese check and fix it.")

    try:
        os.system("pkg install golang")
    except Exception as err:
        print("There is a problem", err, "\nPlaese check and fix it")

    if err == "":
        print("Install Successful!")
    os.system("go run main.go")

try:
    os.system("go")
except Exception as error:
    print("Go is not installed in your system!")
    system = check_system()
    if system == 1:
        print("Plaese install Golang in your system!")
    elif system == 0:
        install_go_on_linux()
