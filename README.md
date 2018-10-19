"# gowalkwindow" 

# Usage
Send serial comm packets And parse the return infos.

# Pack images
I used https://github.com/a-urth/go-bindata

# Compile
go-bindata assets/...
rsrc.exe -manifest main.manifest -o gowalkwindow.syso -ico .\assets\icons\setport.ico
go build .