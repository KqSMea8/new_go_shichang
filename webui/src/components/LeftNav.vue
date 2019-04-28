<template>
  <div id="app">
   LEFT
  </div>
</template>

<script>
export default {
  name: 'LeftNav',
  data () {
    return {
      admins: [
        ['WX Daily Report', 'people_outline'],
        ['Daily Report', 'people_outline'],
        ['Weekly Report', 'people_outline'],
        ['Monthly Check Report', 'people_outline'],
        ['Custom Report', 'people_outline']
      ],
      items: [
        { title: 'TYPES', page: 'Goods', icon: 'dashboard' },
        { title: 'CUSTOMER', page: 'Vendor', icon: 'dashboard' },
        { title: 'PRICE', page: 'VendorPrice', icon: 'dashboard' },
        { title: 'FUNDING', page: 'Deposit', icon: 'dashboard' },
        { title: 'HISTORY', page: 'FixData', icon: 'dashboard' }
      ],
      right: null
    }
  },
  methods: {
    redirect (item) {
      var tag = '/' + item
      this.$router.push({path: tag})
    }
  }
}
</script>
