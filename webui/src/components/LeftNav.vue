<template>
  <div id="app">
    <v-app id="inspire">
      <v-navigation-drawer app permanent>
        <v-toolbar flat>
          <v-list>
            <v-list-tile>
              <v-list-tile-title class="title">
                REPORT
              </v-list-tile-title>
            </v-list-tile>
          </v-list>
        </v-toolbar>

        <v-divider></v-divider>

        <v-list dense class="pt-0">
           <v-list-tile @click="redirect('')">
          <v-list-tile-action>
            <v-icon>home</v-icon>
          </v-list-tile-action>
          <v-list-tile-title>Create Report</v-list-tile-title>
        </v-list-tile>
          <v-list-group
            prepend-icon="account_circle"
            value="true"
            clipped="true"
          >
            <v-list-tile slot="activator">
              <v-list-tile-title>File Download</v-list-tile-title>
            </v-list-tile>
            <v-list-tile
              v-for="(admin, i) in admins"
              :key="i"
            >
              <v-list-tile-title v-text="admin[0]"></v-list-tile-title>
              <v-list-tile-action>
                <v-icon v-text="admin[1]"></v-icon>
              </v-list-tile-action>
            </v-list-tile>
          </v-list-group>
          <v-list-tile v-for="item in items" :key="item.title" @click="redirect(item.page)">
            <v-list-tile-action>
              <v-icon>{{ item.icon }}</v-icon>
            </v-list-tile-action>
            <v-list-tile-content>
              <v-list-tile-title>{{ item.title }}</v-list-tile-title>
            </v-list-tile-content>
          </v-list-tile>
        </v-list>
      </v-navigation-drawer>
    </v-app>
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
