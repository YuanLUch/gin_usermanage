<template>
  <div>
  <div>
    <div>
      <h2>sql画像</h2>
    </div>

    <div class="container">
      <div class="r1">
        <el-radio-group v-model="radio2">
          <el-radio-button label="DDL语句"></el-radio-button>
          <el-radio-button label="DML语句"></el-radio-button>
        </el-radio-group>
      </div>

      <div class="r2">
        <el-radio-group v-model="radio1">
          <el-radio-button label="图形展示"></el-radio-button>
          <el-radio-button label="表格展示"></el-radio-button>
        </el-radio-group>
      </div>
    </div>

    <div style="margin-bottom: 40px;"></div>
    <div v-if="radio1 === '表格展示' && radio2 === 'DML语句'">
      <el-table :data="sqlTableData" border style="width: 100%">
        <el-table-column prop="name" label="名称" width="250">
        </el-table-column>
        <el-table-column prop="num" label="数量" width="250">
        </el-table-column>
        <el-table-column prop="operate" label="源码查看" v-if="this.dmlData.length !== 0">
          <template slot-scope="scope">
          <el-button type="text" @click="showDmlAll(scope.row)">查看源码</el-button>
        </template>
        </el-table-column>
        <el-table-column prop="operate" label="类型查看">
          <template slot-scope="scope">
          <el-button type="text" @click="showDmlDetails(scope.row)">查看详情</el-button>
        </template>
        </el-table-column>
      </el-table>
    </div>

    <div v-else-if="radio1 === '表格展示' && radio2 === 'DDL语句'">
      <el-table :data="ddlTableData" border style="width: 100%">
        <el-table-column prop="name" label="名称" width="250">
        </el-table-column>
        <el-table-column prop="num" label="数量" width="250">
        </el-table-column>
        <el-table-column prop="operate" label="源码查看" v-if="this.ddlData.length !== 0">
          <template slot-scope="scope">
          <el-button type="text" @click="showDdlAll(scope.row)">查看源码</el-button>
        </template>
        </el-table-column>
        <el-table-column prop="operate" label="类型查看">
          <template slot-scope="scope">
          <el-button type="text" @click="showDetails(scope.row)">查看详情</el-button>
        </template>
        </el-table-column>
      </el-table>
    </div>

    <div v-else-if="radio1 === '图形展示' && radio2 === 'DML语句'">
      <sqlTable></sqlTable>
    </div>

    <div v-else>
      <ddlDraw></ddlDraw>
    </div>

  </div>

    <el-dialog title="DML语句详情" :visible.sync="dmlVisible" :width="'70%'">
        <el-table :data="dmlDetailData">
          <el-table-column property="index" label="序号" width="100"></el-table-column>
          <el-table-column property="stmt" label="对象" width="250"></el-table-column>
          <el-table-column property="type" label="类型" width="100"></el-table-column>
          <el-table-column property="num" label="数量" width="100"></el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog title="DML语句源码" :visible.sync="dmlDialogVisible" :width="'70%'">
        <el-table :data="dmlData">
          <el-table-column property="index" label="序号" width="200px"></el-table-column>
          <el-table-column property="statement" label="DML源码"></el-table-column>
      </el-table>
    </el-dialog>

    <el-dialog title="DDL语句源码" :visible.sync="ddlDialogVisible" :width="'70%'">
        <el-table :data="ddlData">
          
          <el-table-column property="statement" label="DDL源码" width="400"></el-table-column>
          <el-table-column property="index" label="序号" width="100"></el-table-column>
      </el-table>
    </el-dialog>

    <!-- <el-dialog title="名称查看" :visible.sync="ddlVisible">
        <el-table :data="ddlDetailData">
          <el-table-column property="index" label="序号" width="100"></el-table-column>
          <el-table-column property="statement" label="名称" width="400"></el-table-column>
      </el-table>
    </el-dialog> -->
    <el-dialog title="DDL语句详情" :visible.sync="ddlVisible" :width="'70%'">
        <el-table :data="ddlDetailData">
          <el-table-column property="index" label="序号" width="100"></el-table-column>
          <el-table-column property="stmt" label="对象" width="250"></el-table-column>
          <el-table-column property="type" label="类型" width="100"></el-table-column>
          <el-table-column property="num" label="数量" width="100"></el-table-column>
      </el-table>
      <!-- <template>
        <div v-if="ddlVisible" id="pieChart" style="width: 100%; height: 400px;"></div>
      </template> -->
    </el-dialog>
    </div>
</template>
  
<script>
// import { Form } from 'element-ui';
import sqlTable from '@/views/sqltable/index.vue'
import ddlDraw from '@/views/ddldraw/index.vue'
import * as echarts from "echarts"

export default {
  name: 'SqlIndex',
  data() {
    return {
      radio1: '图形展示',
      radio2: 'DDL语句',
      sqlTableData: [{
        name: 'table',
        num: 10,
        // operate: '查看'
      }, {
        name: 'view',
        num: 20,
        // operate: '查看'
      }, {
        name: 'trigger',
        num: 30,
        // operate: '查看'
      }],
      ddlTableData: [{
        name: '数据类型',
        num: 10,
        // operate: '查看'
      }, {
        name: '约束',
        num: 20,
        // operate: '查看'
      }, {
        name: '索引',
        num: 30,
        // operate: '查看'
      }],
      dmlVisible: false,
      ddlVisible: false,
      dmlDialogVisible: false,
      ddlDialogVisible: false,
      dmlDetailData: [{
        stmt: '',
        type:'',
        num: 0
      }],
      dmlData: [{
        index: 0,
        statement: '',
      }],
      ddlData: [{
        index: 0,
        statement: '',
      }],
      ddlDetailData: [{
        stmt: '',
        type:'',
        num: 0
      }],
    }
  },

  created() {
    const taskId = this.$route.params.taskId
    console.log("任务ID为", taskId)
    const taskData = this.$store.getters['task/getTaskData'](taskId)
    console.log("拿到的数据为：", taskData)
    const DMLStats = taskData.DMLStats
    const DDLStats = taskData.DDLStats
    const DMLData = []
    const DDLData = []

    for (var key in DMLStats) {
      if (Object.hasOwnProperty.call(DMLStats, key)) {
        var value = DMLStats[key]
        DMLData.push({name: key, num: value})
      }
    }
    console.log("dml数据为：",DMLData)
    this.sqlTableData = DMLData
    
    for (var key in DDLStats) {
      if (Object.hasOwnProperty.call(DDLStats, key)) {
        var value = DDLStats[key]
        DDLData.push({name: key, num: value})
      }
    }
    console.log("ddl数据为：",DDLData)
    this.ddlTableData = DDLData

    const taskType = this.$route.params.taskType
    console.log("taskType", taskType)
    if (taskType === 'Oracle SQL画像') {
        Object.defineProperty(this.dmlData, 'length', { value: 0, writable: false })
        Object.defineProperty(this.ddlData, 'length', { value: 0, writable: false })
      } 
  }, 

  components: { sqlTable, ddlDraw },

  methods: {
    showDmlAll(e) {
        this.dmlDialogVisible = true
        const taskId = this.$route.params.taskId
        const taskData = this.$store.getters['task/getTaskData'](taskId)
        const DMLText = taskData.DMLText
        const DmlData = []
        for (var key in DMLText) {
          if (e.name === key) {
            for (var key_2 in DMLText[key]) {
              DmlData.push({ index: parseInt(key_2) + 1,statement: DMLText[key][key_2] })
            }
          }
        }
        this.dmlData = DmlData
      // console.log(this.dmlData)
    },
    showDmlDetails(e) {
      this.dmlVisible = true
      const taskId = this.$route.params.taskId
      const taskType=this.$route.params.taskType
      const taskData = this.$store.getters['task/getTaskData'](taskId)
      if(taskType=="SQL画像 MySQL"){
        const DMLDetail = taskData.DMLDetail
        const DmlDetailData = []
        var index=1
        for (var key in DMLDetail) {
          if (e.name === key){
            for (var key_2 in DMLDetail[key]['Nodes']) {
              DmlDetailData.push({index:index,stmt: key_2, num: DMLDetail[key]['Nodes'][key_2]})
            }
          }
        }
        this.dmlDetailData = DmlDetailData
      }
      else{
        const DMLDetail=taskData.OracleDMLStats
        const dmldata=[]
        for(var key in DMLDetail){
          if (e.name === key){
            for(var k in DMLDetail[key]){
              if(k!=":"){
                dmldata.push({stmt: k,type:DMLDetail[key][k].sqltype,num: DMLDetail[key][k].num})
              }
            }
          }
        }
        dmldata.sort(function(a,b){
          return a.type.localeCompare(b.type,'zh');
        })
        dmldata.forEach(function(item, index) {
            item.index = index + 1;
        });
        this.dmlDetailData=dmldata
        // console.log("test1:",dmldata)
        // console.log("test2:",this.dmlDetailData)
      }
      
    },
    showDdlAll(e) {
      this.ddlDialogVisible = true
      const taskId = this.$route.params.taskId
      const taskData = this.$store.getters['task/getTaskData'](taskId)
      const DDLText = taskData.DDLText
      const DdlData = []
      for (var key in DDLText) {
        if (e.name === key) {
          for (var key_2 in DDLText[key]) {
            DdlData.push({ index: parseInt(key_2) + 1,statement: DDLText[key][key_2] })
          }
        }
      }
      this.ddlData = DdlData
    },
    showDetails(e) {
      this.ddlVisible = true
      const taskId = this.$route.params.taskId
      const taskType=this.$route.params.taskType
      const taskData = this.$store.getters['task/getTaskData'](taskId)
      if(taskType=="SQL画像 MySQL"){
        const DDLDetail = taskData.DDLDetail
        const data = []
        for (var key in DDLDetail) {
          if (e.name === key){
            for (var key_2 in DDLDetail[key]) {
              data.push({ index: parseInt(key_2) + 1,stmt: DDLDetail[key][key_2] })
            }
          }
        }
        this.ddlDetailData = data
      }
      else{
        const DDLDetail=taskData.OracleDDLStats
        const ddldata=[]
        var index=1
        for(var key in DDLDetail){
          if (e.name === key){
            for(var k in DDLDetail[key]){
              ddldata.push({index:index,stmt: k,type:DDLDetail[key][k].sqltype,num: DDLDetail[key][k].num})
            }
          }
        }
        ddldata.sort(function(a,b){
          return a.type.localeCompare(b.type,'zh');
        })
        ddldata.forEach(function(item, index) {
            item.index = index + 1;
        });
        this.ddlDetailData=ddldata
      }
      
    },
    }
  }
  
</script>

<style>
.container {
  display: flex;
  justify-content: space-between;
}

/* .r1 .r2{
  background-color: aliceblue;
} */
</style>
