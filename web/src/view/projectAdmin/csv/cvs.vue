<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="版本控制器类型" prop="region">
         <el-input v-model="CVSname"></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增版本控制器</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="版本控制器" prop="CVSname" width="120"></el-table-column> 
    
    <el-table-column label="缓存路径" prop="CVSCachePath" width="120"></el-table-column> 
    
    <el-table-column label="编译路径" prop="CVSBuildPath" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateCvs(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteCvs(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="新增版本控制器">
      <el-form ref="elForm" :model="formData" :rules="rules" size="mini" label-width="100px"
        label-position="left">
        <el-row gutter="15">
          <el-form-item label="版本控制器" prop="CVSname">
            <el-input v-model="formData.CVSname" placeholder="git svn ...." clearable
              :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-row>
        <el-row gutter="15">
          <el-form-item label="缓存路径" prop="CVSCachePath">
            <el-input v-model="formData.CVSCachePath" placeholder="maven cache目录" clearable
              :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-row>
        <el-row gutter="15">
          <el-form-item label="编译路径" prop="CVSBuildPath">
            <el-input v-model="formData.CVSBuildPath" placeholder="build目录" clearable
              :style="{width: '100%'}">
            </el-input>
          </el-form-item>
        </el-row>
      </el-form>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createCvs,
    deleteCvs,
    deleteCvsByIds,
    updateCvs,
    findCvs,
    getCvsList
} from "@/api/cvs";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Cvs",
  mixins: [infoList],
  data() {
    return {
      listApi: getCvsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        CVSname:null,CVSCachePath:null,CVSBuildPath:null,
      },
      rules: {
               CVSname: [{
                 required: true,
                 message: '请输入单行文本',
                 trigger: 'blur'
               }],
               CVSCachePath: [{
                 required: true,
                 message: '请输入单行文本',
                 trigger: 'blur'
               }],
               CVSBuildPath: [{
                 required: true,
                 message: '请输入单行文本',
                 trigger: 'blur'
               }],
             },
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteCvsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateCvs(row) {
      const res = await findCvs({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.recsv;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          CVSname:null,
          CVSCachePath:null,
          CVSBuildPath:null,
      };
    },
    async deleteCvs(row) {
      this.visible = false;
      const res = await deleteCvs({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createCvs(this.formData);
          break;
        case "update":
          res = await updateCvs(this.formData);
          break;
        default:
          res = await createCvs(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>