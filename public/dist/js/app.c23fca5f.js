(function(e){function t(t){for(var a,n,s=t[0],i=t[1],r=t[2],b=0,u=[];b<s.length;b++)n=s[b],Object.prototype.hasOwnProperty.call(l,n)&&l[n]&&u.push(l[n][0]),l[n]=0;for(a in i)Object.prototype.hasOwnProperty.call(i,a)&&(e[a]=i[a]);d&&d(t);while(u.length)u.shift()();return o.push.apply(o,r||[]),c()}function c(){for(var e,t=0;t<o.length;t++){for(var c=o[t],a=!0,n=1;n<c.length;n++){var s=c[n];0!==l[s]&&(a=!1)}a&&(o.splice(t--,1),e=i(i.s=c[0]))}return e}var a={},n={app:0},l={app:0},o=[];function s(e){return i.p+"js/"+({}[e]||e)+"."+{"chunk-0c43ac70":"c0d3ae6d","chunk-19619c20":"3fde582f","chunk-27ecac72":"7d0f5595","chunk-280acf90":"0c323559","chunk-41981764":"bc002bb2","chunk-47aa8a59":"fbd59568","chunk-51f84c68":"3868bd88","chunk-655b92da":"e5716eab","chunk-6a13f200":"af069a84","chunk-8bbc52f8":"5f79b4da","chunk-8ca9e7a8":"f308813c","chunk-943314a4":"0acf2fe5","chunk-ab818670":"7e867a59","chunk-ac26e018":"dbe57424","chunk-cbc816ec":"cc303de0","chunk-f752c8ee":"6282ba50"}[e]+".js"}function i(t){if(a[t])return a[t].exports;var c=a[t]={i:t,l:!1,exports:{}};return e[t].call(c.exports,c,c.exports,i),c.l=!0,c.exports}i.e=function(e){var t=[],c={"chunk-0c43ac70":1,"chunk-19619c20":1,"chunk-27ecac72":1,"chunk-280acf90":1,"chunk-41981764":1,"chunk-47aa8a59":1,"chunk-51f84c68":1,"chunk-655b92da":1,"chunk-6a13f200":1,"chunk-8bbc52f8":1,"chunk-8ca9e7a8":1,"chunk-943314a4":1,"chunk-ab818670":1,"chunk-ac26e018":1,"chunk-cbc816ec":1,"chunk-f752c8ee":1};n[e]?t.push(n[e]):0!==n[e]&&c[e]&&t.push(n[e]=new Promise((function(t,c){for(var a="css/"+({}[e]||e)+"."+{"chunk-0c43ac70":"1c979c49","chunk-19619c20":"c4c7fe4f","chunk-27ecac72":"b4c913f3","chunk-280acf90":"69d528ef","chunk-41981764":"3e013bdb","chunk-47aa8a59":"66db8aa2","chunk-51f84c68":"f94effc6","chunk-655b92da":"2c3c6ee7","chunk-6a13f200":"8479488a","chunk-8bbc52f8":"f3c54589","chunk-8ca9e7a8":"bf8a5927","chunk-943314a4":"f6cf2fc8","chunk-ab818670":"42c37833","chunk-ac26e018":"c41c7b2a","chunk-cbc816ec":"00543bcb","chunk-f752c8ee":"7b4f53ca"}[e]+".css",l=i.p+a,o=document.getElementsByTagName("link"),s=0;s<o.length;s++){var r=o[s],b=r.getAttribute("data-href")||r.getAttribute("href");if("stylesheet"===r.rel&&(b===a||b===l))return t()}var u=document.getElementsByTagName("style");for(s=0;s<u.length;s++){r=u[s],b=r.getAttribute("data-href");if(b===a||b===l)return t()}var d=document.createElement("link");d.rel="stylesheet",d.type="text/css",d.onload=t,d.onerror=function(t){var a=t&&t.target&&t.target.src||l,o=new Error("Loading CSS chunk "+e+" failed.\n("+a+")");o.code="CSS_CHUNK_LOAD_FAILED",o.request=a,delete n[e],d.parentNode.removeChild(d),c(o)},d.href=l;var h=document.getElementsByTagName("head")[0];h.appendChild(d)})).then((function(){n[e]=0})));var a=l[e];if(0!==a)if(a)t.push(a[2]);else{var o=new Promise((function(t,c){a=l[e]=[t,c]}));t.push(a[2]=o);var r,b=document.createElement("script");b.charset="utf-8",b.timeout=120,i.nc&&b.setAttribute("nonce",i.nc),b.src=s(e);var u=new Error;r=function(t){b.onerror=b.onload=null,clearTimeout(d);var c=l[e];if(0!==c){if(c){var a=t&&("load"===t.type?"missing":t.type),n=t&&t.target&&t.target.src;u.message="Loading chunk "+e+" failed.\n("+a+": "+n+")",u.name="ChunkLoadError",u.type=a,u.request=n,c[1](u)}l[e]=void 0}};var d=setTimeout((function(){r({type:"timeout",target:b})}),12e4);b.onerror=b.onload=r,document.head.appendChild(b)}return Promise.all(t)},i.m=e,i.c=a,i.d=function(e,t,c){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:c})},i.r=function(e){"undefined"!==typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"===typeof e&&e&&e.__esModule)return e;var c=Object.create(null);if(i.r(c),Object.defineProperty(c,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)i.d(c,a,function(t){return e[t]}.bind(null,a));return c},i.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/",i.oe=function(e){throw console.error(e),e};var r=window["webpackJsonp"]=window["webpackJsonp"]||[],b=r.push.bind(r);r.push=t,r=r.slice();for(var u=0;u<r.length;u++)t(r[u]);var d=b;o.push([0,"chunk-vendors"]),c()})({0:function(e,t,c){e.exports=c("56d7")},5341:function(e,t,c){},"56d7":function(e,t,c){"use strict";c.r(t);var a=c("bc3a"),n=c.n(a),l=c("5333"),o=c("7a23"),s=c("130e"),i=c("2b27"),r=c.n(i);const b={key:0,class:"load"},u={class:"header-content"},d=Object(o["l"])("i",{class:"bx bx-menu"},null,-1),h={key:0,class:"bx bx-sun"},O={key:1,class:"bx bx-moon"},p=Object(o["l"])("i",{class:"bx bx-search"},null,-1),j={class:"sider-item"},f=Object(o["l"])("div",{class:"sider-item-title"},"个人中心",-1),m={class:"navigation"},g={class:"nav-links"},k=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-home"})],-1),v=Object(o["l"])("span",{class:"title"},"主页",-1),y=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-heart"})],-1),x=Object(o["l"])("span",{class:"title"},"最爱",-1),w=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-star"})],-1),_=Object(o["l"])("span",{class:"title"},"收藏",-1),A=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-detail"})],-1),S=Object(o["l"])("span",{class:"title"},"已播放",-1),I={class:"sider-item gallery-list"},R=Object(o["l"])("div",{class:"sider-item-title"},"媒体库",-1),C={class:"navigation more"},L={class:"nav-links"},M={key:0,class:"icon"},P=Object(o["l"])("i",{class:"bx bxs-movie"},null,-1),U=[P],$={key:1,class:"icon"},T=Object(o["l"])("i",{class:"bx bx-desktop"},null,-1),q=[T],E=["data-id"],z={class:"sider-item"},N=Object(o["l"])("div",{class:"sider-item-title"},"管理",-1),D={class:"navigation"},B={class:"nav-links"},F=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bxs-grid"})],-1),H=Object(o["l"])("span",{class:"title"},"媒体中心",-1),J=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-duplicate"})],-1),V=Object(o["l"])("span",{class:"title"},"元数据",-1),K={class:"sider-item"},Q=Object(o["l"])("div",{class:"sider-item-title"},"用户",-1),X={class:"navigation"},Y={class:"nav-links"},G=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-user"})],-1),W=Object(o["l"])("span",{class:"title"},"用户管理",-1),Z=Object(o["l"])("span",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-cog"})],-1),ee=Object(o["l"])("span",{class:"title"},"系统设置",-1),te=Object(o["l"])("i",{class:"bx bx-x"},null,-1),ce=Object(o["l"])("i",{class:"bx bx-search"},null,-1),ae={key:1};function ne(e,t,c,a,n,l){const s=Object(o["R"])("n-button"),i=Object(o["R"])("n-space"),r=Object(o["R"])("n-avatar"),P=Object(o["R"])("n-layout-header"),T=Object(o["R"])("router-link"),ne=Object(o["R"])("n-layout-sider"),le=Object(o["R"])("router-view"),oe=Object(o["R"])("n-layout"),se=Object(o["R"])("n-layout-footer"),ie=Object(o["R"])("n-input"),re=Object(o["R"])("n-card"),be=Object(o["R"])("n-modal"),ue=Object(o["R"])("Login"),de=Object(o["R"])("n-dialog-provider"),he=Object(o["R"])("n-notification-provider"),Oe=Object(o["R"])("n-message-provider"),pe=Object(o["R"])("n-config-provider");return e.load?(Object(o["I"])(),Object(o["k"])("div",b)):(Object(o["I"])(),Object(o["i"])(pe,{key:1,"preflight-style-disabled":"true",theme:e.theme},{default:Object(o["bb"])(()=>[Object(o["n"])(Oe,null,{default:Object(o["bb"])(()=>[Object(o["n"])(he,null,{default:Object(o["bb"])(()=>[Object(o["n"])(de,null,{default:Object(o["bb"])(()=>[e.login?(Object(o["I"])(),Object(o["i"])(oe,{key:0,class:Object(o["y"])([e.dark?"dark":"light","home"])},{default:Object(o["bb"])(()=>[Object(o["n"])(P,{bordered:""},{default:Object(o["bb"])(()=>[Object(o["l"])("div",u,[Object(o["n"])(i,null,{default:Object(o["bb"])(()=>[Object(o["l"])("div",{onClick:t[0]||(t[0]=(...t)=>e.toggDrawer&&e.toggDrawer(...t))},[Object(o["n"])(s,{circle:""},{default:Object(o["bb"])(()=>[d]),_:1})]),Object(o["l"])("div",{onClick:t[1]||(t[1]=(...t)=>e.Home&&e.Home(...t)),class:"title"},Object(o["T"])(e.title),1)]),_:1}),Object(o["n"])(i,{justify:"end"},{default:Object(o["bb"])(()=>[Object(o["n"])(s,{quaternary:"",onClick:t[2]||(t[2]=t=>e.toggDark()),circle:""},{icon:Object(o["bb"])(()=>[e.dark?(Object(o["I"])(),Object(o["k"])("i",h)):(Object(o["I"])(),Object(o["k"])("i",O))]),_:1}),Object(o["n"])(s,{onClick:t[3]||(t[3]=t=>e.showSaerch=!e.showSaerch),circle:""},{icon:Object(o["bb"])(()=>[p]),_:1}),Object(o["n"])(r,{onClick:t[4]||(t[4]=t=>e.LoginOut()),circle:"",size:"medium",src:"https://wework.qpic.cn/wwpic/622138_d-QTzJ_oQAyVDjO_1656146831/0"})]),_:1})])]),_:1}),Object(o["n"])(oe,{position:"absolute",style:{top:"60px",bottom:"60px"},"has-sider":""},{default:Object(o["bb"])(()=>[Object(o["n"])(ne,{collapsed:e.collapsed,"collapse-mode":"width","collapsed-width":0,width:240,"native-scrollbar":!1,bordered:""},{default:Object(o["bb"])(()=>[Object(o["l"])("div",j,[f,Object(o["l"])("div",m,[Object(o["l"])("ul",g,[Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/"},{default:Object(o["bb"])(()=>[k,v]),_:1})]),Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/heart"},{default:Object(o["bb"])(()=>[y,x]),_:1})]),Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/star"},{default:Object(o["bb"])(()=>[w,_]),_:1})]),Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/played"},{default:Object(o["bb"])(()=>[A,S]),_:1})])])])]),Object(o["l"])("div",I,[R,Object(o["l"])("div",C,[Object(o["l"])("ul",L,[(Object(o["I"])(!0),Object(o["k"])(o["b"],null,Object(o["P"])(e.data,(e,t)=>(Object(o["I"])(),Object(o["k"])("li",{key:t},[Object(o["n"])(T,{to:{path:"/list",query:{gallery_uid:e.gallery_uid,gallery_type:e.gallery_type}}},{default:Object(o["bb"])(()=>["movie"==e.gallery_type?(Object(o["I"])(),Object(o["k"])("span",M,U)):(Object(o["I"])(),Object(o["k"])("span",$,q)),Object(o["l"])("span",{"data-id":e.gallery_uid,class:"title"},Object(o["T"])(e.title),9,E)]),_:2},1032,["to"])]))),128))])])]),Object(o["cb"])(Object(o["l"])("div",z,[N,Object(o["l"])("div",D,[Object(o["l"])("ul",B,[Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/gallerys"},{default:Object(o["bb"])(()=>[F,H]),_:1})]),Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/console"},{default:Object(o["bb"])(()=>[J,V]),_:1})])])])],512),[[o["Y"],e.is_admin]]),Object(o["cb"])(Object(o["l"])("div",K,[Q,Object(o["l"])("div",X,[Object(o["l"])("ul",Y,[Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/users"},{default:Object(o["bb"])(()=>[G,W]),_:1})]),Object(o["l"])("li",null,[Object(o["n"])(T,{to:"/setting"},{default:Object(o["bb"])(()=>[Z,ee]),_:1})])])])],512),[[o["Y"],e.is_admin]])]),_:1},8,["collapsed"]),Object(o["n"])(oe,{"native-scrollbar":!1},{default:Object(o["bb"])(()=>[Object(o["n"])(le,{onRefApp:t[5]||(t[5]=t=>e.RefAppData())})]),_:1})]),_:1}),Object(o["n"])(se,{position:"absolute",style:{height:"64px",padding:"24px"},bordered:""},{default:Object(o["bb"])(()=>[Object(o["m"])(" @2022 ")]),_:1}),Object(o["n"])(be,{show:e.showSaerch,"onUpdate:show":t[9]||(t[9]=t=>e.showSaerch=t),"transform-origin":"center"},{default:Object(o["bb"])(()=>[Object(o["n"])(re,{style:{width:"600px"},title:"搜索",bordered:!1,size:"huge",role:"dialog","aria-modal":"true"},{"header-extra":Object(o["bb"])(()=>[Object(o["n"])(s,{onClick:t[6]||(t[6]=t=>e.showSaerch=!e.showSaerch),strong:"",secondary:"",circle:""},{default:Object(o["bb"])(()=>[te]),_:1})]),default:Object(o["bb"])(()=>[Object(o["n"])(ie,{onKeyup:t[7]||(t[7]=Object(o["db"])(t=>e.Search(),["enter"])),value:e.q,"onUpdate:value":t[8]||(t[8]=t=>e.q=t),type:"text",size:"large",placeholder:""},{prefix:Object(o["bb"])(()=>[ce]),_:1},8,["value"])]),_:1})]),_:1},8,["show"])]),_:1},8,["class"])):(Object(o["I"])(),Object(o["k"])("div",ae,[Object(o["n"])(ue,{onIsLogin:t[10]||(t[10]=t=>e.LoginUser()),login:e.login,title:e.title},null,8,["login","title"])]))]),_:1})]),_:1})]),_:1})]),_:1},8,["theme"]))}c("14d9");var le=c("8f5d"),oe=c("0d1c"),se=c.n(oe);const ie=e=>(Object(o["L"])("data-v-540a158a"),e=e(),Object(o["J"])(),e),re={class:"container"},be={class:"top"},ue={class:"header"},de={class:"title"},he={class:"desc"},Oe={class:"main"},pe={class:"md-card login-card"},je=ie(()=>Object(o["l"])("div",{class:"md-card-flex"},[Object(o["l"])("div",{class:"md-card-header-text"})],-1)),fe={class:"create-post-from"},me={class:"form-control"},ge=ie(()=>Object(o["l"])("div",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-envelope"})],-1)),ke={class:"form-control"},ve=ie(()=>Object(o["l"])("div",{class:"icon"},[Object(o["l"])("i",{class:"bx bx-key"})],-1)),ye={class:"form-control"};function xe(e,t,c,a,n,l){const s=Object(o["R"])("n-layout-content");return Object(o["I"])(),Object(o["i"])(s,{class:"login-page"},{default:Object(o["bb"])(()=>[Object(o["l"])("div",re,[Object(o["l"])("div",be,[Object(o["l"])("div",ue,[Object(o["l"])("span",de,Object(o["T"])(c.title),1)]),Object(o["l"])("div",he,Object(o["T"])(c.content),1)]),Object(o["l"])("div",Oe,[Object(o["l"])("div",pe,[je,Object(o["l"])("div",fe,[Object(o["l"])("div",me,[ge,Object(o["cb"])(Object(o["l"])("input",{"onUpdate:modelValue":t[0]||(t[0]=e=>n.user.user_email=e),type:"text",name:"email",placeholder:"账号",required:""},null,512),[[o["X"],n.user.user_email]])]),Object(o["l"])("div",ke,[ve,Object(o["cb"])(Object(o["l"])("input",{"onUpdate:modelValue":t[1]||(t[1]=e=>n.user.user_password=e),type:"password",name:"password",placeholder:"密码",required:""},null,512),[[o["X"],n.user.user_password]])]),Object(o["l"])("div",ye,[Object(o["l"])("button",{class:"btn login-btn",onClick:t[2]||(t[2]=e=>l.LoginUser())},"登录")])])])])])]),_:1})}var we={name:"Login",data(){return{msg:null,user:{user_email:"",user_password:""}}},props:{title:{type:String,default:"Mini Pro"},content:{type:String,default:"一个简洁，好用的私人影库！"},login:{type:Boolean,default:!1}},methods:{RegistUser(){this.axios.post(this.COMMON.apiUrl+"/v1/api/user/create",this.user).then((function(e){200==e.data.code?se.a.show({pos:"top-center",text:"注册成功！",showAction:!1}):se.a.show({pos:"top-center",text:e.data.msg,showAction:!1})}))},LoginUser(){let e=this;this.axios.post(this.COMMON.apiUrl+"/v1/api/user/login",this.user,{headers:{"content-type":"application/json"}}).then((function(t){200==t.data.code?(e.$cookies.set("Authorization",t.data.data,604800),e.$cookies.set("UserId",t.data.user.user_id,604800),se.a.show({pos:"top-center",text:"登录成功！",showAction:!1}),setTimeout((function(){e.$emit("is-login")}),2e3)):se.a.show({pos:"top-center",text:t.data.msg,showAction:!1})}))}}},_e=(c("efae"),c("6b0d")),Ae=c.n(_e);const Se=Ae()(we,[["render",xe],["__scopeId","data-v-540a158a"]]);var Ie=Se,Re=Object(o["o"])({name:"App",components:{Login:Ie},setup(){const e=Object(o["O"])(!1),t=Object(o["O"])(null),c=Object(o["O"])(null),a=Object(o["O"])(!1),n=/Android|webOS|iPhone|iPad|iPod|BlackBerry/i.test(navigator.userAgent);n&&(a.value=!0);const l=Object(o["O"])(!0),s=Object(o["O"])(!1),i=Object(o["O"])(!1),r=Object(o["O"])(!1),b=Object(o["O"])(!1),u=Object(o["O"])(!1),d=Object(o["O"])(null),{proxy:h}=Object(o["p"])();d.value=h.COMMON.title,document.title=d.value;const O=h.$cookies.get("dark"),p=h.$cookies.get("collapsed");function j(){let e=document.querySelectorAll(".sider-item a");e.forEach(t=>{e.forEach(e=>{e.classList.remove("active")}),t.addEventListener("click",()=>{t.classList.add("active")})})}function f(){h.axios.post(h.COMMON.apiUrl+"/v1/api/config/data",{},{headers:{"content-type":"application/json",Authorization:h.$cookies.get("Authorization")}}).then(e=>{if(200==e.data.code){let t=e.data.data;localStorage.setItem("title",t.title),localStorage.setItem("img_url",t.img_url),m()}else se.a.show({pos:"top-center",text:e.data.msg,showAction:!1})}).catch(e=>{se.a.show({pos:"top-center",text:e,showAction:!1}),l.value=!1})}function m(){h.axios.post(h.COMMON.apiUrl+"/v1/api/gallery/list",{},{headers:{"content-type":"application/json",Authorization:h.$cookies.get("Authorization")}}).then(e=>{200==e.data.code?(c.value=e.data.data,setTimeout(()=>{j()},1500),l.value=!1):se.a.show({pos:"top-center",text:e.data.msg,showAction:!1})}).catch(e=>{se.a.show({pos:"top-center",text:e,showAction:!1}),l.value=!1})}function g(){null!=h.$cookies.get("Authorization")?h.axios.get(h.COMMON.apiUrl+"/v1/api/user/data",{headers:{"content-type":"application/json",Authorization:h.$cookies.get("Authorization")}}).then(e=>{200==e.data.code?(s.value=!0,i.value=e.data.data.is_admin,h.$cookies.set("is_admin",i.value,604800),f()):se.a.show({pos:"top-center",text:e.data.msg,showAction:!1})}).catch(e=>{se.a.show({pos:"top-center",text:"登录已过期,请重新登录!",showAction:!1}),console.log(e),l.value=!1}):l.value=!1}"true"==O&&(e.value=!0,t.value=le["a"]),"true"==p&&(a.value=!0);const k=async()=>{g()},v=async()=>{m()};return Object(o["F"])(()=>{g()}),{dark:e,collapsed:a,title:d,load:l,is_admin:i,showIcon:b,showDrawer:r,login:s,darkTheme:le["a"],theme:t,data:c,showSaerch:u,q:Object(o["O"])(null),reF:k,reFApp:v}},methods:{Search(){this.$router.push({path:"/search",query:{q:this.q}})},Home(){this.$router.push({path:"/"})},toggDrawer(){this.collapsed=!this.collapsed,this.$cookies.set("collapsed",this.collapsed)},toggDark(){if(null==this.theme)return this.theme=le["a"],this.$cookies.set("dark","true",2592e3),void(this.dark=!0);"dark"==this.theme.name?(this.theme=null,this.$cookies.set("dark","false",2592e3),this.dark=!1):(this.theme=le["a"],this.$cookies.set("dark","true",2592e3),this.dark=!0)},LoginUser(){this.login=!this.login,this.reF()},LoginOut(){this.$cookies.remove("Authorization"),this.login=!1},RefAppData(){this.reFApp()}}});c("bd52");const Ce=Ae()(Re,[["render",ne]]);var Le=Ce,Me=c("6605");const Pe=[{path:"/",name:"Home",component:()=>c.e("chunk-41981764").then(c.bind(null,"587e"))},{path:"/gallerys",name:"Gallery",component:()=>c.e("chunk-ab818670").then(c.bind(null,"5bfe"))},{path:"/gallerys/works",name:"Work",component:()=>c.e("chunk-f752c8ee").then(c.bind(null,"4088"))},{path:"/gallerys/works/errfiles",name:"Errfile",component:()=>c.e("chunk-655b92da").then(c.bind(null,"3e4b"))},{path:"/list",name:"video_list",component:()=>c.e("chunk-6a13f200").then(c.bind(null,"4bd5"))},{path:"/video",name:"video",component:()=>c.e("chunk-0c43ac70").then(c.bind(null,"85bc"))},{path:"/season",name:"season",component:()=>c.e("chunk-ac26e018").then(c.bind(null,"4bfe"))},{path:"/player",name:"player",component:()=>c.e("chunk-19619c20").then(c.bind(null,"55f4"))},{path:"/person",name:"person",component:()=>c.e("chunk-47aa8a59").then(c.bind(null,"037a"))},{path:"/heart",name:"heart",component:()=>c.e("chunk-8ca9e7a8").then(c.bind(null,"e0c5"))},{path:"/star",name:"star",component:()=>c.e("chunk-8bbc52f8").then(c.bind(null,"9d2a"))},{path:"/played",name:"played",component:()=>c.e("chunk-27ecac72").then(c.bind(null,"d5c9"))},{path:"/users",name:"users",component:()=>c.e("chunk-943314a4").then(c.bind(null,"e382"))},{path:"/search",name:"search",component:()=>c.e("chunk-280acf90").then(c.bind(null,"3ea1"))},{path:"/console",name:"console",component:()=>c.e("chunk-cbc816ec").then(c.bind(null,"bfa9"))},{path:"/setting",name:"setting",component:()=>c.e("chunk-51f84c68").then(c.bind(null,"7424"))}],Ue=Object(Me["a"])({history:Object(Me["b"])(),routes:Pe,scrollBehavior(e,t,c){return{x:0,y:0}}});Ue.afterEach((e,t,c)=>{window.scrollTo(0,0)});var $e=Ue;c("616f");let Te="OneList",qe="",Ee="https://image.tmdb.org";const ze=/Android|webOS|iPhone|iPad|iPod|BlackBerry/i.test(navigator.userAgent);function Ne(){null!=localStorage.getItem("title")&&(Te=localStorage.getItem("title")),null!=localStorage.getItem("img_url")&&(Ee=localStorage.getItem("img_url"))}Ne();var De={apiUrl:qe,title:Te,isMo:ze,imgUrl:Ee};const Be=De;var Fe=Be;c("6da3"),c("678e");const He=Object(o["h"])(Le);He.config.globalProperties.$cookies=r.a,He.config.globalProperties.COMMON=Fe,He.use(l["a"]).use($e).use(s["a"],n.a),He.mount("#app")},b9e5:function(e,t,c){},bd52:function(e,t,c){"use strict";c("5341")},efae:function(e,t,c){"use strict";c("b9e5")}});
//# sourceMappingURL=app.c23fca5f.js.map