// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  data: {
    phone: "",
    rst: "",
    selectTimeStamp: "",
    TimeStamp: ["1","2","4","8"],
    date: '',
    time: '',
    stamp: "间距小时",
    dataList: null,
    motto: 'Hello World',
    userInfo: {},
    hasUserInfo: false,
    canIUse: wx.canIUse('button.open-type.getUserInfo'),
    canIUseGetUserProfile: false,
    canIUseOpenData: wx.canIUse('open-data.type.userAvatarUrl') && wx.canIUse('open-data.type.userNickName') // 如需尝试获取用户信息可改为false
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs',
    })
  },
  FlashTable(phone: string, rst: string, startDate: string, startTime: string, timeStamp: string) {
    wx.request({
      url: `http://192.168.10.97:8080/show?phone=${phone}&rst=${rst}&start_data=${startDate}&start_time=${startTime}&time_stamp=${timeStamp}`,
      method: 'GET',
      header: {
        'content-type': 'application/json'
      },
      success: (res) => {
        console.log(res.data)
        
        this.setData({
          dataList: res.data.data // 将响应数据保存到页面的 data 中
        })
        
      },
      fail: (error) => {
        console.log(error)
      }
    })
  },
  onLoad() {
    this.FlashTable(this.data.phone, this.data.rst, this.data.date, this.data.time, this.data.selectTimeStamp)
  },
  onShow() {
    this.FlashTable(this.data.phone, this.data.rst, this.data.date, this.data.time, this.data.selectTimeStamp)
  },
  onTabItemTap() {
    this.onLoad()
  },
  onTimeChange(e) {
    this.setData({
      time: e.detail.value
    }, this.onLoad)
  },
  onDateChange(e) {
    this.setData({
      date: e.detail.value
    }, this.onLoad)
  },
  onPickerChange(e) {
    console.log("e", e)
    this.setData({
      selectTimeStamp: this.data.TimeStamp[parseInt(e.detail.value)]
    }, this.onLoad)
  },
  SearchPhone(e) {
    this.setData({
      phone: e.detail.value
    }, this.onLoad)
  },
  SearchRst(e) {
    this.setData({
      rst: e.detail.value
    }, this.onLoad)
  }
})
