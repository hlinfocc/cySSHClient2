import{d as h,b as g,B as y,H as i,I as m,l,J as e,R as r,N as B,V as c,u as s,X as u,a7 as b,a8 as I}from"./pinia-5d85bb97.js";import{T as d,B as p,_ as C}from"./index-2c4da577.js";import{T as F}from"./index-43265234.js";import{u as N}from"./vue-94d06e22.js";import"./echarts4-5f68a5d7.js";import"./echarts-0ad49557.js";import"./axios-4d564c32.js";import"./mitt-f7ef348c.js";import"./chevron-up-99364078.js";const S="/static/assets/assets/error-8553e43d.png",k=n=>(b("data-v-3543871e"),n=n(),I(),n),T={class:"container"},V={class:"content"},$={class:"content-main"},x={key:0,class:"result-alert"},E=k(()=>e("img",{src:S,alt:"error"},null,-1)),w={class:"result-btn"},R={class:"result-line"},D=h({__name:"index",setup(n){const{t}=N(),a=g(4);function _(){a.value<4?(a.value+=1,d.message({message:t("menu.result.messageSuccess"),status:"success"})):(a.value=4,d.message({message:t("menu.result.messageEnd"),status:"success"}))}function v(){a.value===4&&(a.value=0)}return(o,H)=>{const f=y("Breadcrumb");return i(),m("div",T,[l(f,{items:["menu.result","menu.result.error"]}),e("div",V,[e("div",$,[a.value===4?(i(),m("div",x,[E,e("div",null,[e("div",null,r(o.$t("menu.result.messageError")),1),e("div",null,r(o.$t("error.result.title")),1)])])):B("",!0),e("div",w,[l(s(p),{type:"primary","native-type":"submit",onClick:_},{default:c(()=>[u(r(o.$t("error.result.home")),1)]),_:1}),l(s(p),{onClick:v},{default:c(()=>[u(r(o.$t("menu.btn.cancel")),1)]),_:1})]),e("div",R,[e("div",null,r(o.$t("menu.line.process")),1),e("div",null,[l(s(F),{data:[{name:s(t)("stepForm.start.coaching")},{name:s(t)("stepForm.immediate.supervisor")},{name:s(t)("stepForm.overall.goals")},{name:s(t)("stepForm.overall.summary")},{name:s(t)("stepForm.overall.end")}],active:a.value,type:"normal"},null,8,["data","active"])])])])])])}}});const G=C(D,[["__scopeId","data-v-3543871e"]]);export{G as default};
