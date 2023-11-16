<template>
  <button class="button button1" text @click="dialogFormVisible = true">发布</button>
  <div>

    <el-dialog v-model="dialogFormVisible" title="添加全国计算机专业大学排名">
      <el-form :model="form">
        <el-form-item label="序号" :label-width="formLabelWidth">
          <el-input v-model="form.index" autocomplete="off" />
        </el-form-item>

        <el-form-item label="学校代号" :label-width="formLabelWidth">
          <el-input v-model="form.code" autocomplete="off" />
        </el-form-item>

        <el-form-item label="学校名称" :label-width="formLabelWidth">
          <el-input v-model="form.name" autocomplete="off" />
        </el-form-item>

        <el-form-item label="等级" :label-width="formLabelWidth">
          <el-select v-model="form.grade" placeholder="请选择一个标签">
            <el-option label="A+" value="A+" />
            <el-option label="A" value="A" />
            <el-option label="A-" value="A-" />
            <el-option label="B+" value="B+" />
            <el-option label="B" value="B" />
            <el-option label="B-" value="B-" />
          </el-select>
        </el-form-item>
      </el-form>


      <template #footer>
        <span class="dialog-footer">
          <el-button @click="dialogFormVisible = false">Cancel</el-button>
          <el-button type="primary" @click="add">
            Confirm
          </el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>


<script>
import { reactive, ref } from 'vue'
import {bus} from '../mitt.js'
import { ElButton, ElDialog, ElTable, ElTableColumn, ElForm, ElFormItem, ElInput, ElSelect, ElOption } from 'element-plus'
export default {
  components: {
    ElButton,
    ElDialog,
    ElTable,
    ElTableColumn,
    ElForm,
    ElFormItem,
    ElInput,
    ElSelect,
    ElOption
  },
  setup() {
    const dialogTableVisible = ref(false)
    const dialogFormVisible = ref(false)
    const formLabelWidth = '140px'

    const form = reactive({
      name: '',
      code: '',
      grade:'',
      index:'',
    })

    return {
      dialogTableVisible,
      dialogFormVisible,
      formLabelWidth,
      form,
    }
  },
  methods:{
    add(){
      this.dialogFormVisible = false
      const data = {
        name: this.form.name,
        code: this.form.code,
        grade: this.form.grade,
        index: this.form.index,
      };
      bus.emit('sendData',data)
      console.log(data)
    },
  }
}

</script>



<style scoped>
.button{
  border: none;
  padding: 10px 50px;
  font-size: 16px;
  border-radius: 25px;
  -webkit-transition-duration: 0.4s; /* Safari */
  transition-duration: 0.4s;
  cursor: pointer;
  width: 270px;
}

.button1 {
  background-color: orangered;
  color: white;
  border: 2px solid orangered;
}

.button1:hover {
  background-color: red;
  border: 2px solid red;
  color: lightyellow;
}

.dialog-footer button:first-child {
  margin-right: 10px;
}

</style>