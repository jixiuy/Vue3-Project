<template>

  <el-scrollbar height="280px">
    <div class="scrollbar-demo-item">
      <p class="aver">序号</p>
      <p class="aver">学校代码</p>
      <p class="aver">学校名称</p>
      <p class="aver">评选结果</p>
      <p class="aver">是否删除</p>
    </div>


    <ul v-infinite-scroll="load" class="infinite-list" style="overflow: auto">
      <div v-for="university in universities" class="infinite-list-item">
          <div style="display: flex;flex-direction: row;justify-content: center;align-items: center;width: 100%">
                <p class="aver" >{{ university.college_rank }}</p>
                <p class="aver">{{ university.code }}</p>
                <p class="aver">{{ university.name }}</p>
                <p class="aver"> {{university.grade}}</p>
                <p class="aver" style="font-size: 15px;color: dodgerblue" @click="deleteItem(university.college_rank)">点击</p>
          </div>
      </div>
    </ul>

  </el-scrollbar>

</template>

<script >
import axios from 'axios';
import {bus} from '../mitt.js'
import { ref } from 'vue'
export default {
  data() {
    return {
      universities: [],
      collegeData:[]
    }
  },setup(){

    const count = ref(100)
    const load = () => {
      count.value += 2
    }
    return{
      count,
      load
    }
  },

  async beforeUpdate() {
    await bus.on('sendData', (data) => {
      if (!this.universities.some((item) => item.code === data.code)) {
        this.universities.push({
          college_rank: data.index,
          code: data.code,
          name: data.name,
          grade: data.grade
        });
        this.collegeData = {
          college_rank: parseInt(data.index),
          code: parseInt(data.code),
          name: data.name,
          grade: data.grade
        };
      }

    })
    const apiUrl = 'http://localhost:8080/colleges'; // 后端API的URL
    // 准备要推送的数据
    // 发送POST请求
    fetch(apiUrl, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(this.collegeData)
    })
        .then(response => response.json())
        .then(data => {
          console.log('Data added successfully:', data);
          // 在这里处理成功响应后的逻辑
        })
        .catch(error => {
          console.error('Error adding data:', error);
          // 在这里处理错误情况
        });
  },
  mounted() {
    axios.get('http://localhost:8080/colleges')
        .then(response => {
          this.universities = response.data;
        })

        .catch(error => {
        //  console.error('获取数据失败:', error);
        });

  },
  methods:{
    deleteItem(rank){
      console.log(this.universities)
      axios.delete(`http://localhost:8080/colleges/${parseInt(rank)}`)
          .then(response => {
            // 处理成功删除的情况
            // console.log('College deleted successfully');
            const indexToDelete = this.universities.findIndex(college => college.college_rank === parseInt(rank));

            if (indexToDelete !== -1) {
              // 使用 splice 方法从数组中删除指定索引的元素
              this.universities.splice(indexToDelete, 1);
            }

          })
          .catch(error => {
            // 处理错误情况
            // console.error(error);
          });
    }
  }
}
</script>


<style scoped>

.aver {
  flex-direction: row;
  text-align: center;
  flex: 20%;
  align-items: center;
  justify-items: center;
}

.scrollbar-demo-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 45px;

  margin-right: 10px;
  margin-top: 10px;
  text-align: center;
  border-radius: 4px;
  background: var(--el-color-primary-light-9);
  color: var(--el-color-primary);

}
.example-pagination-block + .example-pagination-block {
  margin-top: 10px;
}
.example-pagination-block .example-demonstration {
  margin-bottom: 16px;
}
.infinite-list {
  height: 100%;
  padding: 0;
  margin: 0;
  list-style: none;
}
.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 50px;
  background: var(--el-color-primary-light-9);
  margin-right: 10px;
  margin-top: 10px;
  color: var(--el-color-primary);
}
.infinite-list .infinite-list-item + .list-item {

}
</style>