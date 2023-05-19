// index.ts
// 获取应用实例
const app = getApp<IAppOption>()
Page({
  data: {
    isPopupOpen: false,
    formData: {
      token: "",
      name: '',
      licensePlate: '',
      phoneNumber: '',
      departureDate: null,
      departureTime: null,
      departurePlace: '',
      destination: '',
      arrivalDate: null,
      arrivalTime: null,
      totalSeats: '',
      message: ''
    }
  },

  onFocus: function(event) {
    // 获取 <textarea> 标签距离窗口顶部的距离
    wx.createSelectorQuery().select('#textarea').boundingClientRect(function(rect){
      // 将页面滚动到 <textarea> 标签所在的位置
      wx.pageScrollTo({
        scrollTop: rect.top
      })
    }).exec()
  },

  handleInputChange(event) {
    const { field } = event.currentTarget.dataset;
    this.setData({
      [`formData.${field}`]: event.detail.value
    });
  },

  handleSubmit() {
    console.log(this.data.formData)
    for (let key in this.data.formData) {
      console.log(key, "->", this.data.formData[key])
      if (this.data.formData[key] === "" || this.data.formData[key] === null) {
        wx.showToast({
          title: '还有内容没填',
          icon: 'error',
        })
        return
      }
      
    }
    wx.request({
      url: `http://192.168.10.97:8080/new_car`,
      method: 'POST',
      header: {
        'content-type': 'application/json'
      },
      data: {
        token: this.data.formData.token,
        name: this.data.formData.name,
        start_date: this.data.formData.departureDate,
        start_time: this.data.formData.departureTime,
        end_date: this.data.formData.arrivalDate, 
        end_time: this.data.formData.arrivalTime,
        rst: this.data.formData.departurePlace,
        dst: this.data.formData.destination,
        phone: this.data.formData.phoneNumber,
        total: this.data.formData.totalSeats,
        car_number: this.data.formData.licensePlate,
        message: this.data.formData.message,
      },
      success: (res) => {
        console.log("new success", res)
        if (res.data.code === 200) {
          wx.showToast({
            title: '创建成功',
            icon: 'success',
          })
        } else {
          wx.showToast({
            title: res.data.message,
            icon: 'error',
          })
        }
        
      },
      fail: (err) => {
        console.log("new fail", err)
        wx.showToast({
          title: res.data.message,
          icon: 'error',
        })
      }
    })
  }
  
});