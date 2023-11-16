<template>
  <div style="display: flex;align-items: center;flex-direction: column">
    <MyButton></MyButton>

    <div class="box1"  @click="swapColor($event,'box1')">
      <span class="material-icons" :style="{ color: iconColor }">home</span>
      <label :style="{marginLeft:'20px',color:iconColor}" >首页</label>
    </div>

    <div class="box1" @click="swapColor($event,'box2')">
      <span class="material-icons" :style="{ color: iconColor2 }">settings</span>
      <label :style="{marginLeft:'20px',color:iconColor2}" >设置</label>
    </div>

  </div>
</template>

<script>
import MyButton from '../components/MyButton.vue'
export default {
  components:{
    MyButton
  },
  data(){
    return{
      iconColor:"orangered",
      labelColor:"lightgrey",
      iconColor2:"grey",
      labelColor2:"black",
      data:{
        name: '',
        code: '',
        grade: '',
        index: '',
      },
    }
  },
  beforeMount() {
    this.goToHome()
  },

  methods:{
    swapColor(event,parameter){
      if(parameter==='box1'&&this.iconColor==='orangered')return;
      if(parameter==='box2'&&this.iconColor2==='orangered')return;
      if (this.iconColor === "grey") {
        this.iconColor = "orangered"
        this.labelColor = "orangered"
        this.iconColor2 = "grey"
        this.labelColor2 = "black"
        this.goToHome()
      }else if(this.iconColor2 === "grey"){
        this.iconColor2 = "orangered"
        this.labelColor2 = "orangered"
        this.iconColor = "grey"
        this.labelColor = "black"
        this.goToSetting()
      }
    },goToHome() {
      this.$router.push('/home');
    },goToSetting() {
      this.$router.push('/setting');
    }
  }
}

</script>

<style scoped>


.box1{
  display: flex;
  align-items: center;
  margin-top: 20px;
  padding: 10px;
  width: 100px;
}
.box1:hover{
  display: flex;
  align-items: center;
  margin-top: 20px;
  padding: 10px 50px;
  background-color: lightgrey;
  width: 100px;
}

</style>