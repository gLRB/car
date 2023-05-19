// index.ts
// 获取应用实例
const app = getApp<IAppOption>()

Page({
  data: {
    InitData: {},
    EditData: {},
    showModal: false, // 是否显示模态框
  },
  // 事件处理函数
  bindViewTap() {
    wx.navigateTo({
      url: '../logs/logs',
    })
  },
  onLoad(options) {
    console.log("options", options)
    this.setData({
      InitData: options,
      EditData: options,
    })
  },

  showEditModal() {
    this.setData({
      showModal: true,
    })
  },
  hideEditModal() {
    this.setData({
      showModal: false,
    })
  },
  handleFormSubmit() {
    console.log("EditData", this.data.EditData)
    wx.request({
      url: `http://192.168.10.97:8080/edit`,
      method: 'POST',
      header: {
        'content-type': 'application/json'
      },
      data: {
        token: this.data.EditData.token,
        id: this.data.EditData.id,
        name: this.data.EditData.name,
        start_date: this.data.EditData.start_date,
        start_time: this.data.EditData.start_time,
        end_date: this.data.EditData.end_date,
        end_time: this.data.EditData.end_time,
        rst: this.data.EditData.rst,
        dst: this.data.EditData.dst,
        phone: this.data.EditData.phone,
        total: this.data.EditData.total,
        remain: this.data.EditData.remain,
        car_number: this.data.EditData.car_number,
        message: this.data.EditData.message,
      },
      success: (res) => {
        console.log("success", res)
        if (res.data.code === 200) {
          wx.showToast({
            title: '更新成功',
            icon: 'success',
          })
          this.setData({
            InitData: res.data.data,
            EditData: res.data.data,
          })
          this.hideEditModal();
        } else {
          wx.showToast({
            title: res.data.message,
            icon: 'error',
          })
        }
      },
      fail: (err) => {
        console.log("fail", err)
        wx.showToast({
          title: "更新失败",
          icon: 'error',
        })
      }
    })
  },

  handleInputChange(event) {
    const { field } = event.currentTarget.dataset;
    console.log("filed", field)
    this.setData({
      [`EditData.${field}`]: event.detail.value
    });
  },

})
