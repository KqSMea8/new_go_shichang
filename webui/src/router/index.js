import Vue from 'vue'
import Router from 'vue-router'
import Upload from '@/components/Upload'
import Vendor from '@/components/Vendor'
import VendorPrice from '@/components/VendorPrice'
import Goods from '@/components/Goods'
import Vuetify from 'vuetify'
import zhHans from 'vuetify/es5/locale/zh-Hans'

Vue.use(Vuetify, {
  lang: {
    locales: { 'zh-Hans': zhHans },
    current: 'zh-Hans'
  }
})

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Upload',
      component: Upload
    },
    {
      path: '/Vendor',
      name: 'Vendor',
      component: Vendor
    },
    {
      path: '/VendorPrice',
      name: 'VendorPrice',
      component: VendorPrice
    },
    {
      path: '/Goods',
      name: 'Goods',
      component: Goods
    }
  ]
})
