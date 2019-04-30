<template>
   <div id="app">
  <v-app id="inspire">
    <div>
      <v-toolbar flat color="white">
        <v-toolbar-title>客户价格维护</v-toolbar-title>
        <v-divider
          class="mx-2"
          inset
          vertical
        ></v-divider>
        <v-card-title>
        <v-spacer></v-spacer>
        <v-text-field
          v-model="search"
          append-icon="search"
          label="搜索过滤"
          single-line
          hide-details
        ></v-text-field>
      </v-card-title>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="600px">
          <v-btn slot="activator" color="error" dark class="mb-2">批量更新价格</v-btn>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>
            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :disabled="true" v-model="editedItem.Vendor" label="客户名称"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.Price05" label="0-5价格"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.Price12" label="1-2价格"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.Price13" label="1-3价格"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.Price24" label="2-4价格"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.PriceSf" label="石粉"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md3>
                    <v-text-field :rules="dinamicRules" v-model="editedItem.PriceTps" label="头破石"></v-text-field>
                  </v-flex>
                </v-layout>
              </v-container>
            </v-card-text>
            <v-card-actions>
              <v-spacer></v-spacer>
              <v-btn color="blue darken-1" flat @click="close">取消</v-btn>
              <v-btn color="blue darken-1" flat @click="save">保存数据</v-btn>
            </v-card-actions>
          </v-card>
        </v-dialog>
      </v-toolbar>
      <v-alert
        :value="errMessage"
        type="error"
      >
      {{errMessage}}
      </v-alert>
      <v-data-table
        v-model="selected"
        :headers="headers"
        :items="desserts"
        :rows-per-page-items="rowsPerPageItems"
        :search="search"
        item-key="ID"
        select-all
        :pagination.sync="pagination"
        class="elevation-1"
      >
        <template slot="items" slot-scope="props">
            <td class="text-xs-left" style="width:25px">
            <v-checkbox
                v-model="props.selected"
                primary
                hide-details
            ></v-checkbox>
            </td>
          <td class="text-xs-left">{{ props.item.Vendor }}</td>
          <td class="text-xs-left">{{ props.item.Price05 }}</td>
          <td class="text-xs-left">{{ props.item.Price12 }}</td>
          <td class="text-xs-left">{{ props.item.Price13 }}</td>
          <td class="text-xs-left">{{ props.item.Price24 }}</td>
          <td class="text-xs-left">{{ props.item.PriceSf }}</td>
          <td class="text-xs-left">{{ props.item.PriceTps }}</td>
          <td class="justify-left">
            <v-icon small class="mr-2" @click="editItem(props.item)">
              edit
            </v-icon>
          </td>
        </template>
        <template slot="no-data">
          <span>没有价格数据</span>
        </template>
      </v-data-table>
      <div class="text-xs-center pt-2">
        <v-pagination v-model="pagination.page" :length="pages"></v-pagination>
      </div>
    </div>
  </v-app>
</div>
</template>

<script>
export default {
  name: 'VendorPrice',
  data: () => ({
    dialog: false,
    rowsPerPageItems: [ 10, 15, 25, 100, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
    search: '',
    errMessage: '',
    pagination: {},
    selected: [],
    // https://www.cnblogs.com/xinwusuo/p/5948908.html
    dinamicRules: [
      v => /^\d+(\.\d+)?$/.test(v) || '错误的数字格式'
    ],
    requireRules: [
      v => !!v || '客户名称不能为空'
    ],
    headers: [
      { text: '客户名称', width: 160, value: 'Vendor' },
      { text: '0-5价格', width: 80, value: 'Price05' },
      { text: '1-2价格', width: 80, value: 'Price12' },
      { text: '1-3价格', width: 80, value: 'Price13' },
      { text: '2-4价格', width: 80, value: 'Price24' },
      { text: '石粉', width: 80, value: 'PriceSf' },
      { text: '头破石', width: 80, value: 'PriceTps' },
      { text: '操作', value: 'ID', sortable: false }
    ],
    desserts: [],
    editedIndex: -1,
    editedItem: {
      ID: '0',
      Vendor: '',
      Price05: '0',
      Price12: '0',
      Price13: '0',
      Price24: '0',
      PriceSf: '0',
      PriceTps: '0'
    },
    defaultItem: {
      ID: '0',
      Vendor: '',
      Price05: '0',
      Price12: '0',
      Price13: '0',
      Price24: '0',
      PriceSf: '0',
      PriceTps: '0'
    }
  }),

  computed: {
    formTitle () {
      return this.editedIndex === -1 ? '批量更新价格[' + this.selected.length + ']个客户' : '更新价格'
    },
    pages () {
      if (this.pagination.rowsPerPage == null ||
        this.pagination.totalItems == null
      ) return 0

      return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
    }
  },

  watch: {
    dialog (val) {
      val || this.close()
    }
  },

  created () {
    this.pagination.rowsPerPage = 10
    this.initialize()
  },
  methods: {
    initialize () {
      var __this = this
      this.axios.get('http://localhost:8080/vendor')
        .then(function (response) {
          console.log(response.data)
          __this.desserts = response.data
          __this.pagination.totalItems = response.data.length
        })
        .catch(function (error) {
          console.log('error:' + error)
        })
    },

    editItem (item) {
      this.editedIndex = this.desserts.indexOf(item)
      this.editedItem = Object.assign({}, item)
      this.dialog = true
    },

    close () {
      this.dialog = false
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
      }, 300)
    },
    // 更新价格
    vendor_price (item) {
      var __this = this
      __this.errMessage = ''
      this.axios.post('http://localhost:8080/vendor_price', item)
        .then(function (response) {
          console.log(response.data)
          if (response.data.code === 0 && __this.editedIndex > -1) {
            Object.assign(__this.desserts[__this.editedIndex], __this.editedItem)
          } else {
            __this.errMessage = response.data.msg
          }
        })
        .catch(function (error) {
          __this.errMessage = '更新失败，数据异常'
          console.log(error)
        })
    },

    save () {
      if (this.editedIndex > -1) {
        var prices = []
        prices.push(this.editedItem)
        this.vendor_price(prices)
      } else {
        if (this.selected.length === 0) {
          alert('没有选择行')
          this.close()
          return
        }
        for (var i = 0; i < this.selected.length; i++) {
          var row = this.selected[i]
          row.Price05 = this.editedItem.Price05
          row.Price12 = this.editedItem.Price12
          row.Price13 = this.editedItem.Price13
          row.Price24 = this.editedItem.Price24
          row.PriceSf = this.editedItem.PriceSf
          row.PriceTps = this.editedItem.PriceTps
        }
        this.vendor_price(this.selected)
      }
      this.close()
    }
  }
}
</script>
