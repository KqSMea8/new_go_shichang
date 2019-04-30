<template>
  <div class="page-content-container">
    <el-form  label-width="100px">
      <el-row>
      <el-upload  :action="UploadUrl()"  :on-success="UploadSuccess">
              <v-btn
              color="blue-grey"
              class="white--text"
            >
              选择上传文件
              <v-icon right dark>cloud_upload</v-icon>
            </v-btn>
            </el-upload>

      </el-row>

      <el-row>
        <el-col :span="24">
<v-app id="inspire">
    <v-card>
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
      <v-data-table
        :headers="headers"
        :items="list"
        :rows-per-page-items="rowsPerPageItems"
        :search="search"
        :pagination.sync="pagination"
        class="elevation-1"
      >
        <template slot="items" slot-scope="props">
          <td>{{ props.item.RowIndex }}</td>
          <td class="text-xs-left">{{ props.item.NumberFlag }}</td>
          <td class="text-xs-left">{{ props.item.Date }}</td>
          <td class="text-xs-left">{{ props.item.Time }}</td>
          <td class="text-xs-left">{{ props.item.CarNumber }}</td>
          <td class="text-xs-left">{{ props.item.ShopNumber }}</td>
          <td class="text-xs-left">{{ props.item.GrossWeight }}</td>
          <td class="text-xs-left">{{ props.item.TareWeight }}</td>
          <td class="text-xs-left">{{ props.item.NetWeight }}</td>
          <td class="text-xs-left">
            <v-btn outline small color="indigo">{{ props.item.UnitPrice }}</v-btn>
          </td>
          <td class="text-xs-left">{{ props.item.TakeDept }}</td>
        </template>
      </v-data-table>
      <div class="text-xs-center pt-2">
        <v-pagination v-model="pagination.page" :length="pages"></v-pagination>
      </div>
    </v-card>
  </v-app>
      </el-col>
    </el-row>

    </el-form>
  </div>
</template>

<script>
import Vue from 'vue'
import uploader from 'vue-simple-uploader'
Vue.use(uploader)

export default {
  name: 'Upload',
  data () {
    return {
      list: [],
      rowsPerPageItems: [ 10, 15, 25, 100, { 'text': '$vuetify.dataIterator.rowsPerPageAll', 'value': -1 } ],
      search: '',
      pagination: {},
      selected: [],
      headers: [
        {text: '', width: 40, align: 'left', value: 'RowIndex'},
        { text: '磅单编号', width: 120, value: 'NumberFlag' },
        { text: '首次称重日期', width: 100, value: 'Date' },
        { text: '毛重时间', width: 60, value: 'Time' },
        { text: '车号', width: 120, value: 'CarNumber' },
        { text: '货物名称', width: 80, value: 'ShopNumber' },
        { text: '毛重', width: 80, value: 'GrossWeight' },
        { text: '皮重', width: 80, value: 'TareWeight' },
        { text: '净重', width: 80, value: 'NetWeight' },
        { text: '单价', width: 80, value: 'UnitPrice' },
        { text: '收货单位', value: 'TakeDept' }

      ]
    }
  },
  components: {
  },
  created () {
    this.pagination.rowsPerPage = -1
    // this.post_test()
  },
  computed: {
    pages () {
      if (this.pagination.rowsPerPage == null ||
        this.pagination.totalItems == null
      ) return 0

      return Math.ceil(this.pagination.totalItems / this.pagination.rowsPerPage)
    }
  },
  methods: {
    UploadUrl () {
      return 'http://localhost:8080/upload'
    },
    get_login_info () {
      var __this = this
      this.axios.get('http://localhost:8080/ping')
        .then(function (response) {
          console.log(response)
          __this.msg = response.data
        })
        .catch(function (error) {
          console.log('error:' + error)
        })
    },
    post_test () {
      this.axios.post('http://localhost:8080/pp', null)
        .then(function (response) {
          console.log(1)
          console.log(response)
        })
        .catch(function (error) {
          console.log(error)
        })
    },
    UploadSuccess (rsp) {
      this.list = rsp
      this.pagination.totalItems = rsp.length
    }
  }
}
</script>

<style>
  .uploader-example {
    width: 880px;
    padding: 15px;
    margin: 40px auto 0;
    font-size: 12px;
    box-shadow: 0 0 10px rgba(0, 0, 0, .4);
  }
  .uploader-example .uploader-btn {
    margin-right: 4px;
  }
  .uploader-example .uploader-list {
    max-height: 440px;
    overflow: auto;
    overflow-x: hidden;
    overflow-y: auto;
  }
</style>
