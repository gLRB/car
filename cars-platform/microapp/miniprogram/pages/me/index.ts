// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  // 点击按钮时触发
  onTapGetUserProfile: function() {
    this.setData({
      showModal: true
    })
    wx.getUserProfile({
      desc: '用于获取您的微信昵称', // 向用户说明获取信息的用途
      success: res => {
        console.log("onTapGetUserProfile", res.userInfo) // 打印用户信息
        // 将用户信息保存到全局变量中，方便其他页面使用
        app.globalData.userInfo = res.userInfo
      },
      fail: res => {
        console.log("onTapGetUserProfile fail ", res)
      }
    })
  },
  data: {
    showModal: false,
    dataList: [
      {
        name: "11",
        gender: "nn"
      },
      {
        name: "11",
        gender: "nn"
      },
      {
        name: "11",
        gender: "nn"
      },
      {
        name: "11",
        gender: "nn"
      }
  ],
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
  onLoad() {
    wx.login({
      success: function(res) {
        console.log("res", res);
      }
    });
    this.getUserProfile()

  },
  getUserProfile() {
    // 推荐使用wx.getUserProfile获取用户信息，开发者每次通过该接口获取用户个人信息均需用户确认，开发者妥善保管用户快速填写的头像昵称，避免重复弹窗
    wx.getUserProfile({
      desc: '展示用户信息', // 声明获取用户个人信息后的用途，后续会展示在弹窗中，请谨慎填写
      success: (res) => {
        console.log("getUserProfile", res)
        this.setData({
          userInfo: res.userInfo,
          hasUserInfo: true
        })
      }, 
      fail: (res) => {
        console.log("getUserProfile fail", res)
      }
    })
  },
  getUserInfo(e: any) {
    // 不推荐使用getUserInfo获取用户信息，预计自2021年4月13日起，getUserInfo将不再弹出弹窗，并直接返回匿名的用户个人信息
    console.log(e)
    this.setData({
      userInfo: e.detail.userInfo,
      hasUserInfo: true
    })
  }
})
