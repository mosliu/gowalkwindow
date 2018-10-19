package ui

import (
	"fmt"
	"github.com/lxn/walk"
	"github.com/mosliu/gowalkwindow/serialport"
	"github.com/tarm/serial"
)
import . "github.com/lxn/walk/declarative"

type AppWindow struct {
	*walk.MainWindow
	recentMenu                   *walk.Menu
	outPutTextEdit               *walk.TextEdit
	btnOpenPort, btnFetchVersion *walk.PushButton
	versionShowLineEdit          *walk.LineEdit
	//status bar items
	currentPort serialport.CommPort
}

var defaultPortConfig = serial.Config{
	Name:     "COM1",
	Baud:     9600,
	Parity:   serial.ParityNone,
	StopBits: serial.Stop1,
}

var Appw = AppWindow{
	//当前串口配置信息
	currentPort: serialport.CommPort{
		Config: &defaultPortConfig,
	},
}

//Draw UI
func ShowMainApp() {
	log.Debugln("Show Main Window")
	//Appw := new(AppWindow)
	var btn *walk.PushButton
	//菜单项
	var openAction, showAboutBoxAction, setupAction *walk.Action
	//var recentMenu *walk.Menu

	//图标
	//iconPortSet, err := walk.NewIconFromFile("./assets/icons/setport.ico")
	//if err != nil {
	//    log.Fatal(err)
	//}

	icon1, err := walk.NewIconFromFile("./assets/imgs/check.ico")
	if err != nil {
		logf.Fatal(err)
	}
	icon2, err := walk.NewIconFromFile("./assets/imgs/stop.ico")
	if err != nil {
		logf.Fatal(err)
	}
	var sbi *walk.StatusBarItem
	err = MainWindow{
		AssignTo: &Appw.MainWindow,
		Name:     "haliluya",
		Title:    "兰光版本号下位机读取工具",
		MinSize:  Size{800, 600},
		//渐变色背景
		Background: GradientBrush{
			//定义点 0-4 5个点 分别在左上 右上 正中 右下 左下
			Vertexes: []walk.GradientVertex{
				{X: 0, Y: 0, Color: walk.RGB(255, 255, 127)},
				{X: 1, Y: 0, Color: walk.RGB(127, 191, 255)},
				{X: 0.5, Y: 0.5, Color: walk.RGB(255, 255, 255)},
				{X: 1, Y: 1, Color: walk.RGB(127, 255, 127)},
				{X: 0, Y: 1, Color: walk.RGB(255, 127, 127)},
			},
			//使用上面定义的点，绘制渐变三角形
			Triangles: []walk.GradientTriangle{
				{0, 1, 2},
				{1, 3, 2},
				{3, 4, 2},
				{4, 0, 2},
			},
		},
		Layout: VBox{
			MarginsZero: true,
		},
		MenuItems: []MenuItem{
			Menu{
				Text: "&File",
				Items: []MenuItem{
					Action{
						AssignTo:    &openAction,
						Text:        "&Open",
						Image:       "./assets/imgs/open.png",
						Enabled:     Bind("enabledCB.Checked"),
						Visible:     Bind("!openHiddenCB.Checked"),
						Shortcut:    Shortcut{walk.ModControl, walk.KeyO},
						OnTriggered: Appw.openAction_Triggered,
					},
					Menu{
						AssignTo: &Appw.recentMenu,
						Text:     "Recent",
					},
					Separator{},
					Action{
						Text:        "E&xit",
						OnTriggered: func() { Appw.Close() },
					},
				},
			},
			Menu{
				Text: "&Setup",
				Items: []MenuItem{
					Action{
						Image:       "./assets/icons/setport.ico",
						AssignTo:    &setupAction,
						Text:        "Setup",
						OnTriggered: Appw.setupAction_Triggered,
					},
				},
			},
			Menu{
				Text: "&Help",
				Items: []MenuItem{
					Action{
						AssignTo:    &showAboutBoxAction,
						Text:        "About",
						OnTriggered: Appw.showAboutBoxAction_Triggered,
					},
				},
			},
		},
		StatusBarItems: []StatusBarItem{
			StatusBarItem{
				AssignTo: &sbi,
				Icon:     icon1,
				Text:     "click",
				Width:    80,
				OnClicked: func() {
					if sbi.Text() == "click" {
						sbi.SetText("again")
						sbi.SetIcon(icon2)
					} else {
						sbi.SetText("click")
						sbi.SetIcon(icon1)
					}
				},
			},
			StatusBarItem{
				Text:        "left",
				ToolTipText: "no tooltip for me",
			},
			StatusBarItem{
				Text: "\tcenter",
			},
			StatusBarItem{
				Text: "\t\tright",
			},
			StatusBarItem{
				Icon:        icon1,
				ToolTipText: "An icon with a tooltip",
			},
		},
		Children: []Widget{
			VSpacer{},
			//主区域
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Composite{
						MinSize: Size{Width: 250},
						Layout:  Grid{Columns: 2},
						Children: []Widget{
							VSpacer{ColumnSpan: 2},
							PushButton{
								Text:     "打开串口",
								AssignTo: &Appw.btnOpenPort,
								OnClicked: func() {
									btnOpenPortClicked(Appw.btnOpenPort)
								},
							},
							PushButton{
								Text:     "获取版本",
								AssignTo: &Appw.btnFetchVersion,
								OnClicked: func() {
									btnFetchVersionClicked(Appw.btnFetchVersion)
								},
							},
							VSpacer{ColumnSpan: 2},
							LineEdit{
								Font:       Font{PointSize: 20},
								AssignTo:   &Appw.versionShowLineEdit,
								Text:       "Versions：",
								ColumnSpan: 2,
								ReadOnly:   true,
							},
							VSpacer{ColumnSpan: 2},
						},
					},
					HSplitter{},
					Composite{
						MinSize: Size{Width: 550},
						Layout:  VBox{},
						Children: []Widget{
							TextEdit{
								AssignTo: &Appw.outPutTextEdit,
								Text:     "",
							},
							PushButton{
								Text:     "push me",
								AssignTo: &btn,
								OnClicked: func() {
									btnClicked(btn)
								},
							},
							Label{
								Text: "点击按钮后，可拖动标题栏测试界面状态。",
							},
						},
					},
				},
			},

			VSpacer{},
		},
	}.Create()
	if err != nil {
		log.Fatal(err)
	}
	Appw.addRecentFileActions("Foo", "Bar", "Baz")
	Appw.Run()
}

func btnClicked(btn *walk.PushButton) {
	btn.SetText("哈哈")
	//start("COM1")
}

//打开串口按钮
func btnOpenPortClicked(btn *walk.PushButton) {

	if btn.Text() == "打开串口" {
		btn.SetText("关闭串口")
		openSerialPort()
	} else {
		btn.SetText("打开串口")
		closeSerialPort()
	}
	//start("COM1")
}

//关闭串口
func closeSerialPort() {
	Appw.DisplayInfo("Closing SerialPort...")
	err := Appw.currentPort.ClosePort()
	if err != nil {
		Appw.DisplayInfo("\r\nSerialPort %v close failed.error:%v", Appw.currentPort.Config.Name, err)
	}
	Appw.DisplayInfo("\r\nSerialPort %v Closed.", Appw.currentPort.Config.Name)
}

//打开串口
func openSerialPort() {
	// 这里使用 log.Debugf的办法 会导致无法显示。很奇怪。
	// 先用fmt.SPrintf转换 在用log.Debug也不行 都会显示不出来。
	// 其他位置显示正常，考虑是否是因为使用了go func的原因？
	// logF可以输出。
	Appw.DisplayInfo("Opening SerialPort...")
	err := Appw.currentPort.OpenPort()
	if err != nil {
		Appw.DisplayInfo("\r\nSerialPort %v open failed.error:%v", Appw.currentPort.Config.Name)
		Appw.btnOpenPort.SetText("打开串口")
	} else {
		Appw.DisplayInfo("\r\nSerialPort %v Opened.", Appw.currentPort.Config.Name)
	}
	//str:= fmt.Sprintf("%d,Port Open \r\n",5474)
	//logrus.Debug(str)
	//log.Debug(str)
	//fmt.Print(str)
	//log.Debugf("%d,Port Open \r\n",5474)
	//log.Infof("%d,Port Open \r\n",5474)
	//log.Warnf("%d,Port Open \r\nfun",5474)
	//log.Debugln("Port Open ",config)
	//logf.Debugf("%d,Port Open \r\n",5474)
	//logf.Infof("%d,Port Open \r\n",5474)
	//logf.Warnf("%d,Port Open \r\n",5474)
	//logf.Debugln("Port Open ",config)
}

//获取版本按键
func btnFetchVersionClicked(btn *walk.PushButton) {

	if btn.Text() == "获取版本" {
		btn.SetText("停止获取版本")
		startFetchVersion()
	} else {
		btn.SetText("获取版本")
		stopFetchVersion()
	}

	//start("COM1")
}

func startFetchVersion() {
	Appw.DisplayInfo("\r\nStart Fetching Device Electric Version")
	Appw.currentPort.StartHandle()
}

func stopFetchVersion() {

}

//菜单open项
func (appw *AppWindow) openAction_Triggered() {
	log.Debug("Open Action Triggered!")
	//walk.MsgBox(Appw, "Open", "Pretend to open a file...", walk.MsgBoxIconInformation)
}

//菜单关于项
func (appw *AppWindow) showAboutBoxAction_Triggered() {
	walk.MsgBox(appw, "About", "Labthink Version Fetch Tool\r\nV1.0.0\r\nAuthor：Moses", walk.MsgBoxIconInformation)
}

//菜单，串口设置项
func (appw *AppWindow) setupAction_Triggered() {
	//Appw.config = serial.Config{
	//    Name:     "COM1",
	//    Baud:     9600,
	//    Parity:   serial.ParityNone,
	//    StopBits: serial.Stop1,
	//}
	if cmd, err := RunSerialPortSetDialog(appw, appw.currentPort.Config); err != nil {
		log.Print(err)
	} else if cmd == walk.DlgCmdOK {
		appw.outPutTextEdit.AppendText(fmt.Sprintf("\r\n%+v", appw.currentPort.Config))
	}
}

//动态添加例子
func (appw *AppWindow) addRecentFileActions(texts ...string) {
	for _, text := range texts {
		a := walk.NewAction()
		a.SetText(text)
		a.Triggered().Attach(appw.openAction_Triggered)
		appw.recentMenu.Actions().Add(a)
	}
}

//从界面中显示提示信息
func (appw *AppWindow) DisplayInfo(format string, args ...interface{}) {
	showstr := format
	if len(args) != 0 {
		showstr = fmt.Sprintf(format, args)
	}
	appw.outPutTextEdit.AppendText(showstr)
	log.Debug(showstr)
}

func (appw *AppWindow) SetVersionShowLineEditText(str string) {
	appw.versionShowLineEdit.SetText(str)
}

//--------------------------------------------------------------------------------------------------------------------
// -------------------------------------------------------------------------------------------------------------------

var done = make(chan bool, 1)

//
//func start(PortNum string) {
//    //serialPort := serialport.OpenPort(PortNum)
//    //serialPort := Appw.currentPort.Port
//    //serialPort := serial.Config{
//    //    Name:PortNum
//    //}
//    recvChan := make(chan []byte)
//    sendChan := make(chan []byte)
//    packetChan := make(chan handlers.Packet)
//
//    var pi handlers.PacketInfo
//    //pi.PreDefine()
//
//    //go serialport.StartSendChannel(serialPort, sendChan)
//
//    //serialConfig := &serial.Config{Name: "COM1", Baud: 115200}
//
//    //
//    //for b := range recvChan {
//    //    fmt.Println(b)
//    //}
//    bs := PrepareSendPacket(0x10)
//    time.AfterFunc(time.Millisecond*100, func() {
//        bs[3] = 1
//        sendChan <- bs
//        time.AfterFunc(time.Millisecond*100, func() {
//            sendChan <- bs
//        })
//    })
//
//    //start listen to port
//    time.AfterFunc(time.Millisecond*1300, func() {
//        //go serialport.StartRecvSerialData(serialPort, recvChan)
//        go serialport.HandleRawData(recvChan, packetChan, pi)
//        go handlePacket(packetChan, sendChan, pi)
//    })
//
//    time.AfterFunc(time.Millisecond*1400, func() {
//        //预读下位机数据
//        time.AfterFunc(time.Millisecond*100, func() {
//            getCurrentDeviceConf(sendChan)
//        })
//        //开始校对时间
//        time.AfterFunc(time.Millisecond*300, func() {
//            startCorrectDevice(sendChan)
//        })
//    })
//
//    //要数据
//    //开始
//
//    //done := make(chan bool, 1)
//    <-done
//    color.Blue("本次校正结束,可以关闭了")
//    var out string
//    fmt.Scanln(&out)
//
//    //fmt.Println("buf",buf)
//    //logs.Printf("%q", buf[:n])
//}
//
//func PrepareSendPacket(funcId byte) []byte {
//    bs := make([]byte, 12)
//    binary.BigEndian.PutUint16(bs[0:], 0xAAAA)
//    bs[2] = funcId //return code
//    binary.BigEndian.PutUint16(bs[10:], 0xBBBB)
//    return bs
//}
//
//type TimeLog struct {
//    SNum  int
//    STime int
//    Time  time.Time
//}
//
//type DeviceStatus struct {
//    THTL uint16
//}
//
//var lowerMachine = DeviceStatus{0}
//var timeLogArray = []TimeLog{}
//var data_source = []float64{}
//var startingFlag = false    //如果设定为true，5s还无回复设定为false，则有问题
//var gettingTHTLFlag = false //如果设定为true，5s还无回复设定为false，则有问题
//var finishedFlag = false    //如果设定为true，5s还无回复设定为false，则有问题
//
//func handlePacket(packetChan chan handlers.Packet, sendChan chan []byte, pi handlers.PacketInfo) {
//    for pkt := range packetChan {
//        body := pkt.Content[len(pi.HeadStyle) : len(pi.HeadStyle)+pi.BodyLength]
//        log.Debug("got packet:", pkt.Content)
//        log.Debugf("body:", body)
//        switch body[0] {
//        case 0x01:
//            {
//            }
//        case 0x02:
//            {
//                //0x02 0x03 = 0x02 0x02: 发送下位机存储的时间校正参数
//                //0x04 0x05: bigEndian 存储的参数
//                gettingTHTLFlag = false
//
//                lowerMachine.THTL = binary.BigEndian.Uint16(body[1:3])
//                log.Info(
//                    color.GreenString(fmt.Sprintf(
//                        "Lower Machine's Th and Tl is:%d [ %X ]", lowerMachine.THTL, lowerMachine.THTL)))
//            }
//        case 0x10:
//            //时间调整接收
//            {
//                startingFlag = false
//                //AA AA 10 00 02 00 00 03 E9 00 BB BB
//                no := int(binary.BigEndian.Uint16(body[1:3]))
//                stime := int(binary.BigEndian.Uint32(body[3:7]))
//                now := time.Now()
//                newTimeLog := TimeLog{no, stime, now}
//                const startCalcNum = 20
//                if no == 0 {
//                    log.Info("Start time Adjust")
//                    log.Info("Reset old logs Array,start new Array")
//                    timeLogArray = []TimeLog{}
//                } else if no < startCalcNum+2 {
//                    log.Info("buffering", strings.Repeat(".", no))
//                } else if no > startCalcNum+1 {
//                    timegoes := now.Sub(timeLogArray[startCalcNum].Time).Nanoseconds()
//                    //timegoes := timeLogArray[len(timeLogArray)-1].Time.Sub(timeLogArray[0].Time).Nanoseconds()
//                    timegoes = timegoes / 1000
//                    stime := newTimeLog.STime - timeLogArray[startCalcNum].STime
//                    diff := (int(timegoes) - stime) / (no - startCalcNum)
//                    log.Infof("Data: [No：%d] [diff:%d us] [remotelogtime:%d] [timePast:%d]", newTimeLog.SNum, diff, stime, timegoes)
//                    if no > 40 {
//                        data_source = append(data_source, float64(diff))
//                    }
//                    if no > 60 {
//                        μ, σ, a := calcGaussianDistribution()
//                        if no > 120 && a > 0.14 {
//                            log.Infof("结束校正！！！。平均数：%f，标准差：%f,分布：%f", μ, σ, a)
//                            //发送停止指令
//
//                            time.AfterFunc(time.Millisecond*50, func() {
//                                bs := PrepareSendPacket(0x10)
//                                bs[3] = 1
//                                sendChan <- bs
//                                time.AfterFunc(time.Millisecond*150, func() {
//                                    sendChan <- bs
//                                })
//                            })
//
//                            time.AfterFunc(time.Millisecond*300, func() {
//                                if finishedFlag == false {
//                                    finishedFlag = true
//                                    bs := PrepareSendPacket(0x02)
//                                    bs[3] = 1 //设定
//                                    toset := uint16(int32(lowerMachine.THTL) + int32(math.Ceil(float64(μ)/float64(41.6))))
//                                    binary.BigEndian.PutUint16(bs[4:6], toset)
//
//                                    log.Info(color.GreenString(fmt.Sprintf("Writing %X to LowerMachine", bs[4:6])))
//                                    sendChan <- bs
//                                    //设置第2次
//                                    time.AfterFunc(time.Millisecond*1000, func() {
//                                        sendChan <- bs
//                                    })
//                                    //设置第3次
//                                    time.AfterFunc(time.Millisecond*2000, func() {
//                                        sendChan <- bs
//                                    })
//                                    //设置第4次
//                                    time.AfterFunc(time.Millisecond*3000, func() {
//                                        sendChan <- bs
//                                    })
//                                    //设置第5次
//                                    time.AfterFunc(time.Millisecond*4000, func() {
//                                        sendChan <- bs
//                                    })
//                                    //重启模块
//                                    time.AfterFunc(time.Millisecond*6000, func() {
//                                        bs = PrepareSendPacket(0x33)
//                                        bs[3] = 0x33 //设定
//                                        sendChan <- bs
//                                    })
//
//                                    time.AfterFunc(time.Millisecond*5000, func() {
//                                        //var logs TimeLog
//                                        log.Debugf("%4s %10s %10s ", "No.", "RemoteTime", "LocalTime")
//                                        for i, alog := range timeLogArray {
//                                            last := i
//                                            if last > 0 {
//                                                last = i - 1
//                                            }
//                                            pastTime := alog.Time.Sub(timeLogArray[last].Time).Nanoseconds() / 1000
//                                            log.Debugf("%4d %10d %10d ", alog.SNum, alog.STime, pastTime)
//                                        }
//                                        done <- true
//                                    })
//                                }
//                            })
//
//                        }
//
//                    }
//                }
//                timeLogArray = append(timeLogArray, newTimeLog)
//            }
//
//        }
//
//    }
//
//}
//
//func getCurrentDeviceConf(sendChan chan []byte) {
//    bs := PrepareSendPacket(0x02)
//    bs[3] = 0x02
//    gettingTHTLFlag = true
//    sendChan <- bs
//    time.AfterFunc(time.Millisecond*5000, func() {
//        if gettingTHTLFlag {
//            getCurrentDeviceConf(sendChan)
//        }
//    })
//}
//
//func startCorrectDevice(sendChan chan []byte) {
//    bs := PrepareSendPacket(0x10)
//    bs[3] = 0
//    startingFlag = true
//    sendChan <- bs
//    time.AfterFunc(time.Millisecond*5000, func() {
//        if startingFlag {
//            startCorrectDevice(sendChan)
//        }
//    })
//}
//
//func calcGaussianDistribution() (μ, σ, a float64) {
//    //数学期望
//    var sum float64 = 0
//    sort.Float64s(data_source)
//    arraySize := len(data_source)
//    arrayCutSize := arraySize / 3
//    array := data_source[arrayCutSize : arraySize-arrayCutSize]
//
//    for _, v := range array {
//        sum += v
//    }
//    //μ := float64(sum) / float64(len(data_source))
//    μ = float64(sum) / float64(len(array))
//    //标准差
//    var variance float64
//    for _, v := range array {
//        variance += math.Pow(v-μ, 2)
//    }
//    σ = math.Sqrt(variance / float64(len(array)))
//    //正态分布公式
//    a = 1 / (math.Sqrt(2*math.Pi) * σ) * math.Pow(math.E, -math.Pow(μ-μ, 2)/(2*math.Pow(σ, 2)))
//    //fmt.Println(a)
//    log.Infof("样品均值：%f，标准差：%f,分布：%f", μ, σ, a)
//    return
//}
