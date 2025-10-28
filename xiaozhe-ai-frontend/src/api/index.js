import axios from 'axios';

const api=axios.create({
    baseURL:'http://localhost:8088/api'
});


const GetUser=function (){
    api.get('/user/getUserInfo').then(
        (res)=>{
            console.log(res.data)
        }
    )
}
GetUser ()