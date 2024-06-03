import{bH as R,bI as h,w as _,ah as P,bJ as C,v as K,aD as L,a4 as A,av as U,aw as H,bK as c,a6 as z,bL as J}from"./index-2c4da577.js";const B=/%[sdj%]/g,N=()=>{};function E(e){if(!e||!e.length)return null;const t={};return e.forEach(n=>{const i=n.field;t[i]=t[i]||[],t[i].push(n)}),t}function b(e,...t){if(typeof e=="function")return e(...t);if(typeof e=="string"){let n=0;const i=t.length;return String(e).replace(B,a=>{if(a==="%%")return"%";if(n>=i)return a;switch(a){case"%j":try{return JSON.stringify(t[n++])}catch{return"[Circular]"}case"%d":return Number(t[n++]);case"%s":return String(t[n++]);default:return a}})}return e}function G(e){return["string","url","hex","email","pattern","digits","time","dateYMD","longDateTime","dateTime","dateYM","version","speczh","specialch","specialch2","acceptImg","acceptFile","fileSize"].includes(e)}function l(e,t){return!!(P(e)||t==="array"&&Array.isArray(e)&&!e.length||G(t)&&typeof e=="string"&&!e)}function W(e){return Object.keys(e).length===0}function Q(e,t,n){let i=0;const s=[],a=e.length;function r(d){s.push(...d),i++,i===a&&n(s)}e.forEach(d=>{t(d,r)})}function Y(e,t,n){let i=0;const s=e.length;function a(r){if(r&&r.length){n(r);return}const d=i;i=i+1,d<s?t(e[d],a):n([])}a([])}function X(e){const t=[];return Object.keys(e).forEach(n=>{t.push(...e[n])}),t}function I(e,t,n,i){if(t.first){const q=new Promise((p,u)=>{const y=u,M=S=>(i(S),S.length?y({errors:S,fields:E(S)}):p()),F=X(e);Y(F,n,M)});return q.catch(p=>p.errors&&p.fields||R.logger.error(p)),q}let s=t.firstFields||[];s===!0&&(s=Object.keys(e));let a=0;const r=Object.keys(e),d=r.length,o=[],f=new Promise((q,p)=>{const u=p,y=M=>{if(o.push(...M),a++,a===d)return i(o),o.length?u({errors:o,fields:E(o)}):q()};r.forEach(M=>{const F=e[M];s.includes(M)?Y(F,n,y):Q(F,n,y)})});return f.catch(q=>q.errors&&q.fields||R.logger.error(q)),f}function $(e){return t=>t&&t.message?(t.field=t.field||e.fullField,t):{message:typeof t=="function"?t():t,field:t.field||e.fullField}}function O(e,t){if(!t)return e;for(const n in t)if(h.call(t,n)){const i=t[n];typeof i=="object"&&typeof e[n]=="object"?e[n]=_(_({},e[n]),i):e[n]=i}return e}const V=Object.freeze(Object.defineProperty({__proto__:null,asyncMap:I,complementError:$,convertFieldsError:E,deepMerge:O,format:b,isEmptyObject:W,isEmptyValue:l,warning:N},Symbol.toStringTag,{value:"Module"}));function m(e,t){m.getSystemMessage=()=>m.getDefaultMessage(t),m.messages=m.getSystemMessage(t),m.systemMessages=m.messages,this.rules=null,this._messages=m.systemMessages,this.define(e)}const k=e=>t=>{let n,i=[],s={};function a(r){Array.isArray(r)?i=i.concat(...r):i.push(r)}for(n=0;n<t.length;n++)a(t[n]);i.length?s=E(i):(i=null,s=null),e(i,s)},ee=(e,t)=>{let n=(e.type==="object"||e.type==="array")&&(typeof e.fields=="object"||typeof e.defaultField=="object");return n=n&&(e.required||!e.required&&t.value),n},te=(e,t)=>{let n={};function i(s,a){return K(_({},a),{fullField:`${e.fullField}.${s}`})}if(e.defaultField)for(const s in t.value)h.call(t.value,s)&&(n[s]=e.defaultField);n=_(_({},n),t.rule.fields);for(const s in n)if(h.call(n,s)){const a=Array.isArray(n[s])?n[s]:[n[s]];n[s]=a.map(i.bind(null,s))}return n},ie=e=>(Array.isArray(e)||(e=[e]),e),ne=({rule:e,failds:t,options:n})=>(e.message?t=[].concat(e.message).map($(e)):n.error?t=[n.error(e,b(n.messages.required,e.field))]:t=[],t),se=({data:e,options:t})=>{if(e.rule.options){let{messages:n,error:i}=t;Object.assign(e.rule.options,{messages:n,error:i})}},ae=({failds:e,doIt:t})=>n=>{const i=[];e&&e.length&&i.push(...e),n&&n.length&&i.push(...n),t(i.length?i:null)},re=(e,t,n,i,s)=>(a=[])=>{let r=a;const d=ee(t,s);if(r=ie(r),!e.suppressWarning&&r.length&&m.warning("async-validator:",r),r.length&&t.message&&(r=[].concat(t.message)),r=r.map($(t)),e.first&&r.length)return n[t.field]=1,i(r);if(d){if(t.required&&!s.value)return r=ne({rule:t,failds:r,options:e}),i(r);const o=new m(te(t,s));o.messages(e.messages),se({data:s,options:e}),o.validate(s.value,s.rule.options||e,ae({failds:r,doIt:i}))}else i(r)};m.prototype={messages(e){return e&&(this._messages=O(m.getSystemMessage(),e)),this._messages},define(e){if(!e)throw new Error("Cannot configure a schema with no rules");if(Array.isArray(e)||typeof e!="object")throw new TypeError("Rules must be an object");this.rules={};let t;Object.keys(e).forEach(n=>{h.call(e,n)&&(t=e[n],this.rules[n]=Array.isArray(t)?t:[t])})},getSeries(e,t,n){let i,s;const a={};return(e.keys||Object.keys(this.rules)).forEach(d=>{i=this.rules[d],s=t[d],i.forEach(o=>{let f=o;typeof f.transform=="function"&&(t===n&&(t=_({},t)),t[d]=f.transform(s),s=t[d]),typeof f=="function"?f={validator:f}:f=_({},f),f.validator=this.getValidationMethod(f),f.field=d,f.fullField=f.fullField||d,f.type=this.getType(f),e.custom&&Object.assign(f,e.custom),f.validator&&(a[d]=a[d]||[],a[d].push({rule:f,value:s,source:t,field:d}))})}),a},mergeMessage(e){if(e.messages){let t=this.messages();t===m.systemMessages&&(t=m.getSystemMessage()),O(t,e.messages),e.messages=t}else e.messages=this.messages()},validate(e,t={},n=()=>{}){let i=e,s=t,a=n;if(typeof s=="function"&&(a=s,s={}),!this.rules||Object.keys(this.rules).length===0)return a&&a(),Promise.resolve();const r=k(a);this.mergeMessage(s);const d=this.getSeries(s,i,e),o={};return I(d,s,(f,q)=>{const p=f.rule,u=re(s,p,o,q,f);let y;p.asyncValidator?y=p.asyncValidator(p,f.value,u,f.source,s):p.validator&&(y=p.validator(p,f.value,u,f.source,s),y===!0?u():y===!1?u(p.message||`${p.field} fails`):Array.isArray(y)?u(y):y instanceof Error&&u(y.message)),y&&y.then&&y.then(()=>u(),M=>u(M))},f=>{r(f)})},getValidationMethod(e){if(C(e.validator))return e.validator;const t=Object.keys(e),n=t.indexOf("message");return n>-1&&t.splice(n,1),t.length===1&&t[0]==="required"?m.validators.required:m.validators[this.getType(e)]||!1},getType(e){if(e.type===void 0&&e.pattern instanceof RegExp&&(e.type="pattern"),typeof e.validator!="function"&&e.type&&!h.call(m.validators,e.type))throw new Error(b("Unknown rule type %s",e.type));return e.type||"string"}};m.register=(e,t)=>{if(typeof t!="function")throw new TypeError("Cannot register a validator by type, validator is not a function");m.validators[e]=t};m.validators={};m.warning=N;m.messages={};m.systemMessages={};m.getDefaultMessage=()=>{};var T=m;function Z({rule:e,checkValue:t,source:n,errors:i,options:s,type:a}){e.required&&(!h.call(n,e.field)||l(t,a||e.type))&&i.push(b(s.messages.required,""))}const fe='^(([^<>()\\[\\]\\\\.,;:\\s@"]+(\\.[^<>()\\[\\]\\\\.,;:\\s@"]+)*)|(".+"))',de=new RegExp(fe+"@((\\[[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}\\.[0-9]{1,3}])|(([a-zA-Z\\-0-9]+\\.)+[a-zA-Z]{2,6}))$"),x={acceptImg:/\.(png|jpe?g|gif)$/,acceptFile:/\.(doc?x|xls?x|ppt?x|txt)$/,email:de,fileSize:/^\d+(\.\d+)?[KMGTPEZY]?B$/i,hex:/^#?([a-f0-9]{6}|[a-f0-9]{3})$/i,speczh:/^[0-9a-zA-Z_\u4E00-\u9FA5]+$/,specialch:/^[0-9a-zA-Z_\-.]+$/,specialch2:/^[0-9a-zA-Z_-]+$/,url:/^(([a-zA-Z]{3,}):)?\/\/([\w-]+\.)+[\w]+(\/[a-zA-Z- ./?%&=]*)?/i,version:/^\d+\.\d+(\.\d+)*$/},j={integer:e=>j.number(e)&&/^[-]?[\d]+$/.test(e),float:e=>j.number(e)&&!j.integer(e),array:Array.isArray,regexp(e){if(e instanceof RegExp)return!0;try{return!!new RegExp(e)}catch{return!1}},date:L,number:e=>A(Number(e)),object:e=>U(e)&&!j.array(e),method:e=>H(e)==="function",email:e=>c(e)||!!e.match(x.email)&&e.length<255,url:e=>c(e)||!!e.match(x.url),hex:e=>c(e)||!!e.match(x.hex),digits:e=>c(e)||/^\d+$/.test(e),time:e=>c(e)||/^((0)[0-9]|1[0-9]|20|21|22|23):([0-5][0-9])$/.test(e),dateYM:e=>c(e)||z(e,"yyyy-MM")===e,dateYMD:e=>c(e)||z(e,"yyyy-MM-dd")===e,dateTime:e=>c(e)||z(e,"yyyy-MM-dd hh:mm")===e,longDateTime:e=>c(e)||z(e,"yyyy-MM-dd hh:mm:ss")===e,version:e=>c(e)||!!e.match(x.version),speczh:e=>c(e)||!!e.match(x.speczh),specialch:e=>c(e)||!!e.match(x.specialch),specialch2:e=>c(e)||!!e.match(x.specialch2),acceptImg:e=>c(e)||!!e.match(x.acceptImg),acceptFile:e=>c(e)||!!e.match(x.acceptFile),fileSize:e=>c(e)||!!e.match(x.fileSize)};function oe(e,t,n,i,s){if(e.required&&t===void 0){Z(e);return}const a=["array","acceptImg","acceptFile","date","digits","dateTime","dateYM","dateYMD","email","float","fileSize","hex","integer","longDateTime","method","number","object","regexp","speczh","specialch","specialch2","time","version","url"],r=e.type;a.includes(r)?j[r](t)||i.push(b(s.messages.types[r],"",e.type)):r&&typeof t!==e.type&&i.push(b(s.messages.types[r],"",e.type))}function ge({min:e,max:t,val:n,key:i,rule:s,errors:a,util:r,options:d}){e&&!t&&n<s.min?a.push(r.format(d.messages[i].min,"",s.min)):t&&!e&&n>s.max?a.push(r.format(d.messages[i].max,"",s.max)):e&&t&&(n<s.min||n>s.max)&&a.push(r.format(d.messages[i].range,"",s.min,s.max))}function me(e,t,n,i,s){const a=A(e.len),r=A(e.min),d=A(e.max);let o=t,f=null;const q=A(Number(t)),p=typeof t=="string",u=Array.isArray(t);if(q?f="number":p?f="string":u&&(f="array"),!f)return!1;u&&(o=t.length),p&&(o=J(t,"string")),e.type==="number"&&(o=t),a?o!==e.len&&i.push(b(s.messages[f].len,"",e.len)):ge({min:r,max:d,val:o,key:f,rule:e,errors:i,util:V,options:s})}const w="enum";function pe(e,t,n,i,s){e[w]=Array.isArray(e[w])?e[w]:[],e[w].includes(t)||i.push(b(s.messages[w],"",e[w].join(", ")))}function le(e,t,n,i,s){e.pattern&&(e.pattern instanceof RegExp?(e.pattern.lastIndex=0,e.pattern.test(t)||i.push(b(s.messages.pattern.mismatch,"",t,e.pattern))):typeof e.pattern=="string"&&(new RegExp(e.pattern).test(t)||i.push(b(s.messages.pattern.mismatch,"",t,e.pattern))))}function ye(e,t,n,i,s){(/^\s+$/.test(t)||t==="")&&i.push(b(s.messages.whitespace,""))}var g={type:oe,range:me,pattern:le,required:Z,whitespace:ye,enum:pe};function ce(e,t,n,i,s){const a=[],r=e.required||!e.required&&h.call(i,e.field),d=o=>o&&typeof o=="string"&&new Date(o).toString()!=="Invalid Date";if(r){if(l(t)&&!e.required)return n();if(g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),!l(t)){let o;typeof t=="number"||d(t)?o=new Date(t):o=t,g.type(e,o,i,a,s),o&&typeof o.getTime=="function"&&g.range(e,o.getTime(),i,a,s)}}n(a)}function v(e,t,n,i,s){const a=e.type,r=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t,a)&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:r,options:s,type:a}),l(t,a)||g.type(e,t,i,r,s)}n(r)}function ve(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t)&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),t!==void 0&&(g.type(e,t,i,a,s),g.range(e,t,i,a,s))}n(a)}function he(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t,"array")&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s,type:"array"}),l(t,"array")||(g.type(e,t,i,a,s),g.range(e,t,i,a,s))}n(a)}function ue(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t,"string")&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s,type:"string"}),l(t,"string")||(g.type(e,t,i,a,s),g.range(e,t,i,a,s),g.pattern(e,t,i,a,s),e.whitespace===!0&&g.whitespace(e,t,i,a,s))}n(a)}function D(e,t,n,i,s){const a=e.required||!e.required&&h.call(i,e.field),r=[];if(a){if(!e.required&&l(t))return n();g.required({rule:e,checkValue:t,source:i,errors:r,options:s}),t!==void 0&&g.type(e,t,i,r,s)}n(r)}function qe(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(t===""&&(t=void 0),!e.required&&l(t))return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),t!==void 0&&(g.type(e,t,i,a,s),g.range(e,t,i,a,s))}n(a)}function be(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t)&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),t!==void 0&&t!==""&&(g.type(e,t,i,a,s),g.range(e,t,i,a,s))}n(a)}function xe(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t,"string")&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),l(t,"string")||g.pattern(e,t,i,a,s)}n(a)}function Me(e,t,n,i,s){const a=[],r=Array.isArray(t)?"array":typeof t;g.required({rule:e,checkValue:t,source:i,errors:a,options:s,type:r}),n(a)}const _e="enum";function we(e,t,n,i,s){const a=[];if(e.required||!e.required&&h.call(i,e.field)){if(l(t)&&!e.required)return n();g.required({rule:e,checkValue:t,source:i,errors:a,options:s}),t!==void 0&&g[_e](e,t,i,a,s)}n(a)}var Ae={date:ce,float:ve,array:he,string:ue,method:D,number:qe,integer:be,pattern:xe,required:Me,hex:v,url:v,time:v,email:v,digits:v,dateYM:v,speczh:v,dateYMD:v,version:v,fileSize:v,regexp:D,object:D,dateTime:v,specialch:v,boolean:D,acceptImg:v,specialch2:v,acceptFile:v,longDateTime:v,enum:we};const je=e=>({string:e("validation.types.string"),method:e("validation.types.method"),array:e("validation.types.array"),object:e("validation.types.object"),number:e("validation.types.number"),date:e("validation.types.date"),boolean:e("validation.types.boolean"),integer:e("validation.types.integer"),float:e("validation.types.float"),regexp:e("validation.types.regexp"),email:e("validation.types.email"),url:e("validation.types.url"),hex:e("validation.types.hex"),digits:e("validation.types.digits"),time:e("validation.types.time"),dateYM:e("validation.types.dateYM"),dateYMD:e("validation.types.dateYMD"),dateTime:e("validation.types.dateTime"),longDateTime:e("validation.types.longDateTime"),version:e("validation.types.version"),speczh:e("validation.types.speczh"),specialch:e("validation.types.specialch"),specialch2:e("validation.types.hex"),acceptImg:e("validation.types.acceptImg"),acceptFile:e("validation.types.acceptFile"),fileSize:e("validation.types.fileSize")});var Fe=(e=t=>t)=>({default:e("validation.default"),required:e("validation.required"),enum:e("validation.enum"),whitespace:e("validation.whitespace"),date:{format:e("validation.date.format"),parse:e("validation.date.parse"),invalid:e("validation.date.invalid")},types:je(e),string:{len:e("validation.string.len"),min:e("validation.string.min"),max:e("validation.string.max"),range:e("validation.string.range")},number:{len:e("validation.number.len"),min:e("validation.number.min"),max:e("validation.number.max"),range:e("validation.number.range")},array:{len:e("validation.array.len"),min:e("validation.array.min"),max:e("validation.array.max"),range:e("validation.array.range")},pattern:{mismatch:e("validation.pattern.mismatch")}});T.validators=Ae;T.getDefaultMessage=Fe;var ze=T;export{ze as v};
