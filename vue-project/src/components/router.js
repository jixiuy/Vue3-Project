import {createRouter,createWebHashHistory} from "vue-router";
import Home from "../components/Home.vue"
import Setting from "../components/Setting.vue"

const router = createRouter({
    history:createWebHashHistory(),
    routes:[
        {
            path:"/home",
            component:Home
        },
        {
            path:"/setting",
            component:Setting
        }
    ]
})

export default router;