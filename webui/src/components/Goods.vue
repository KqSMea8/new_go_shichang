<template>
   <div id="app">
  <v-app id="inspire">
    <div>
      <v-toolbar flat color="white">
        <v-toolbar-title>石粉型号</v-toolbar-title>
        <v-divider
          class="mx-2"
          inset
          vertical
        ></v-divider>
        <v-spacer></v-spacer>
        <v-dialog v-model="dialog" max-width="600px">
          <v-btn slot="activator" :disabled="true" color="primary" dark class="mb-2">新增数据</v-btn>
          <v-card>
            <v-card-title>
              <span class="headline">{{ formTitle }}</span>
            </v-card-title>
            <v-card-text>
              <v-container grid-list-md>
                <v-layout wrap>
                  <v-flex xs12 sm6 md4>
                    <v-text-field :disabled="isEditItem" :rules="requireRules" v-model="editedItem.GoodsName" label="*石粉名称"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md4>
                    <v-text-field mask="###" v-model="editedItem.OrderIndex" label="排序"></v-text-field>
                  </v-flex>
                   <v-flex xs12 sm6 md4>
                    <v-text-field v-model="editedItem.Nodes" label="别名"></v-text-field>
                  </v-flex>
                  <v-flex xs12 sm6 md4>
                    <v-radio-group v-model="editedItem.GoodsType" row>
                        <v-radio label="石仔" value="石仔"></v-radio>
                        <v-radio label="石粉" value="石粉"></v-radio>
                    </v-radio-group>
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
        :headers="headers"
        :items="desserts"
         :rows-per-page-items="rowsPerPageItems"
        :search="search"
        :pagination.sync="pagination"
        class="elevation-1"
      >
        <template slot="items" slot-scope="props">
          <td class="text-xs-left">{{ props.item.GoodsName }}</td>
          <td class="text-xs-left">{{ props.item.GoodsType }}</td>
          <td class="text-xs-left">{{ props.item.OrderIndex }}</td>
          <td class="text-xs-left">{{ props.item.Nodes }}</td>
          <td class="justify-left">
            <v-icon small class="mr-2" @click="editItem(props.item)">
              edit
            </v-icon>
          </td>
        </template>
        <template slot="no-data">
          <span>没有石粉数据</span>
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
  name: 'Goods',
  // inject: ['reload'],
  data: () => ({
    dialog: false,
    rowsPerPageItems: [ 10, 15, 25, 100, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
    search: '',
    isEditItem: false, // 是否新增
    errMessage: '',
    pagination: {},
    selected: [],
    // https://www.cnblogs.com/xinwusuo/p/5948908.html
    dinamicRules: [
      v => /^\d+(\.\d+)?$/.test(v) || '错误的数字格式'
    ],
    requireRules: [
      v => !!v || '石粉名称不能为空'
    ],
    headers: [
      { text: '石粉名称', width: 160, value: 'GoodsName' },
      { text: '类别', width: 160, value: 'GoodsType' },
      { text: '排序', width: 140, value: 'OrderIndex' },
      { text: '别名', width: 120, value: 'Nodes' },
      { text: '操作', value: 'name', sortable: false }
    ],
    desserts: [],
    editedIndex: -1,
    editedItem: {
      GoodsName: '',
      GoodsType: '石仔',
      Nodes: '',
      OrderIndex: 0
    },
    defaultItem: {
      GoodsName: '',
      GoodsType: '石仔',
      Nodes: '',
      OrderIndex: 0
    }
  }),

  computed: {
    formTitle () {
      return this.editedIndex === -1 ? '新增数据' : '编辑数据'
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
      this.axios.get('http://localhost:8080/goods')
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
      this.isEditItem = true
      this.dialog = true
    },

    deleteItem (item) {
      const index = this.desserts.indexOf(item)
      confirm('是否删除数据?') && this.goods_delete(item, index)
    },

    close () {
      this.dialog = false
      setTimeout(() => {
        this.editedItem = Object.assign({}, this.defaultItem)
        this.editedIndex = -1
        this.isEditItem = false
      }, 300)
    },
    // 新增数据
    goods_insert (item) {
      var __this = this
      __this.errMessage = ''
      this.axios.post('http://localhost:8080/goods_insert', item)
        .then(function (response) {
          console.log(response.data)
          if (response.data.code === 0) {
            __this.reload()
          } else {
            __this.errMessage = response.data.msg
          }
        })
        .catch(function (error) {
          __this.errMessage = '新增数据失败，数据异常'
          console.log(error)
        })
    },
    // 删除数据
    goods_delete (item, index) {
      var __this = this
      __this.errMessage = ''
      this.axios.post('http://localhost:8080/goods_delete', item)
        .then(function (response) {
          console.log(response.data)
          if (response.data.code === 0) {
            __this.desserts.splice(index, 1)
            __this.pagination.totalItems = __this.desserts.length
          } else {
            __this.errMessage = response.data.msg
          }
        })
        .catch(function (error) {
          __this.errMessage = '删除数据失败，数据异常'
          console.log(error)
        })
    },
    // 更新数据
    goods_update (item) {
      var __this = this
      __this.errMessage = ''
      this.axios.post('http://localhost:8080/goods_update', item)
        .then(function (response) {
          console.log(response.data)
          if (response.data.code === 0) {
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

    reload () {
      this.initialize()
    },

    save () {
      if (this.editedIndex > -1) {
        this.goods_update(this.editedItem)
      } else {
        this.goods_insert(this.editedItem)
      }
      this.close()
    }
  }
}
</script>
