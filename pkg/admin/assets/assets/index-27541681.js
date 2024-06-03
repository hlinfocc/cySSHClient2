import{d as f,p as C,a as A,$ as H,b as V,c as F,a7 as z,a8 as M,A as W,a9 as J}from"./index-2c4da577.js";import{H as w,I as k,Q as c,B as y,J as d,O as p,a0 as B,$ as m,l as h,X as g,R as $,V as R,W as P,Y as L}from"./pinia-5d85bb97.js";const Q=({emit:e,props:r,state:t})=>n=>{n=[].concat(n);const a=r.accordion?n[0]:n;t.activeNames=n,e("update:modelValue",a),e("change",a)},X=({api:e,props:r,state:t})=>n=>{const a=t.activeNames.slice(0),i=a.indexOf(n==null?void 0:n.name);(()=>{let o=r.beforeClose?r.beforeClose(n,t.activeNames):!0;return new Promise(l=>{o&&o.then?o.then(()=>l(!0)).catch(()=>l(!1)):l(o)})})().then(o=>{r.accordion?(o||!a.length)&&e.setActiveNames(a[0]===(n==null?void 0:n.name)?"":n==null?void 0:n.name):(i>-1?o&&a.splice(i,1):a.push(n==null?void 0:n.name),e.setActiveNames(a))})},D=["state"],U=(e,{reactive:r,watch:t},{parent:n,emit:a,constants:i})=>{const s=i.EVENT_NAME.CollapseItemClick,o=r({activeNames:[]}),l={state:o,setActiveNames:Q({emit:a,props:e,state:o})};return l.handleItemClick=X({api:l,props:e,state:o}),t(()=>e.modelValue,u=>{o.activeNames=u||u===0?[].concat(u):[]},{immediate:!0,deep:!0}),n.$on(s,l.handleItemClick),l};function q(e,r){var t=typeof Symbol<"u"&&e[Symbol.iterator]||e["@@iterator"];if(t)return(t=t.call(e)).next.bind(t);if(Array.isArray(e)||(t=G(e))||r&&e&&typeof e.length=="number"){t&&(e=t);var n=0;return function(){return n>=e.length?{done:!0}:{done:!1,value:e[n++]}}}throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}function G(e,r){if(e){if(typeof e=="string")return O(e,r);var t=Object.prototype.toString.call(e).slice(8,-1);if(t==="Object"&&e.constructor&&(t=e.constructor.name),t==="Map"||t==="Set")return Array.from(e);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return O(e,r)}}function O(e,r){(r==null||r>e.length)&&(r=e.length);for(var t=0,n=new Array(r);t<r;t++)n[t]=e[t];return n}var Z=function(r,t){for(var n=r.__vccOpts||r,a=q(t),i;!(i=a()).done;){var s=i.value,o=s[0],l=s[1];n[o]=l}return n},x=f({props:[].concat(C,["accordion","modelValue","beforeClose"]),setup:function(r,t){return A({props:r,context:t,renderless:U,api:D})}}),ee={class:"tiny-collapse",role:"tablist","aria-multiselectable":"true"};function te(e,r,t,n,a,i){return w(),k("div",ee,[c(e.$slots,"default")])}var N=Z(x,[["render",te]]);function re(e,r){var t=typeof Symbol<"u"&&e[Symbol.iterator]||e["@@iterator"];if(t)return(t=t.call(e)).next.bind(t);if(Array.isArray(e)||(t=ne(e))||r&&e&&typeof e.length=="number"){t&&(e=t);var n=0;return function(){return n>=e.length?{done:!0}:{done:!1,value:e[n++]}}}throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}function ne(e,r){if(e){if(typeof e=="string")return E(e,r);var t=Object.prototype.toString.call(e).slice(8,-1);if(t==="Object"&&e.constructor&&(t=e.constructor.name),t==="Map"||t==="Set")return Array.from(e);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return E(e,r)}}function E(e,r){(r==null||r>e.length)&&(r=e.length);for(var t=0,n=new Array(r);t<r;t++)n[t]=e[t];return n}var ae=function(r,t){for(var n=r.__vccOpts||r,a=re(t),i;!(i=a()).done;){var s=i.value,o=s[0],l=s[1];n[o]=l}return n},oe=f({props:[].concat(C,["accordion","modelValue","beforeClose"]),emits:["update:modelValue","change"],setup:function(r,t){return A({props:r,context:t,renderless:U,api:D})}}),ie={role:"tablist","aria-multiselectable":"true","data-tag":"tiny-collapse"};function se(e,r,t,n,a,i){return w(),k("div",ie,[c(e.$slots,"default")])}var le=ae(oe,[["render",se]]);function I(){return I=Object.assign?Object.assign.bind():function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},I.apply(this,arguments)}var de=function(r){var t,n=typeof process=="object"?(t={})==null?void 0:t.TINY_MODE:null;return(n||r)==="pc"?N:(n||r)==="mobile-first"?le:N},ce={COMPONENT_NAME:{Collapse:"Collapse"},EVENT_NAME:{CollapseItemClick:"collapse-item.click"}},ue=I({},H,{_constants:{type:Object,default:function(){return ce}},accordion:Boolean,beforeClose:Function,modelValue:{type:[Array,String,Number],default:function(){return[]}}}),v=f({name:V+"Collapse",componentName:"Collapse",props:ue,provide:function(){return{collapse:this}},setup:function(r,t){return F({props:r,context:t,template:de})}}),pe="3.15.0";v.model={prop:"modelValue",event:"update:modelValue"};v.install=function(e){e.component(v.name,v)};v.version=pe;const fe=({state:e,interval:r})=>()=>{setTimeout(()=>{e.isClick?e.isClick=!1:e.focusing=!0},r)},ve=({componentName:e,dispatch:r,eventName:t,props:n,parent:a,state:i})=>()=>{n.disabled||(r(e,t,a),i.focusing=!1,i.isClick=!0)},me=({componentName:e,dispatch:r,eventName:t,parent:n})=>()=>r(e,t,n),K=["state","isActive","handleFocus","handleEnterClick","handleHeaderClick"],Y=(e,{computed:r,reactive:t},{parent:n,constants:a,dispatch:i})=>{const s=n.collapse._constants,o=s.COMPONENT_NAME.Collapse,l=s.EVENT_NAME.CollapseItemClick,u=t({id:z(),isClick:!1,focusing:!1,contentHeight:0,contentWrapStyle:{height:"auto",display:"block"},isActive:r(()=>n.collapse.state.activeNames.includes(e.name))});return{state:u,handleFocus:fe({state:u,interval:a.INTERVAL}),handleEnterClick:me({componentName:o,dispatch:i,eventName:l,parent:n}),handleHeaderClick:ve({componentName:o,dispatch:i,eventName:l,props:e,parent:n,state:u})}};function be(e,r){var t=typeof Symbol<"u"&&e[Symbol.iterator]||e["@@iterator"];if(t)return(t=t.call(e)).next.bind(t);if(Array.isArray(e)||(t=ye(e))||r&&e&&typeof e.length=="number"){t&&(e=t);var n=0;return function(){return n>=e.length?{done:!0}:{done:!1,value:e[n++]}}}throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}function ye(e,r){if(e){if(typeof e=="string")return S(e,r);var t=Object.prototype.toString.call(e).slice(8,-1);if(t==="Object"&&e.constructor&&(t=e.constructor.name),t==="Map"||t==="Set")return Array.from(e);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return S(e,r)}}function S(e,r){(r==null||r>e.length)&&(r=e.length);for(var t=0,n=new Array(r);t<r;t++)n[t]=e[t];return n}var he=function(r,t){for(var n=r.__vccOpts||r,a=be(t),i;!(i=a()).done;){var s=i.value,o=s[0],l=s[1];n[o]=l}return n},ge=f({props:[].concat(C,["title","titleRight","name","disabled"]),components:{CollapseTransition:M,IconChevronRight:W()},setup:function(r,t){return A({props:r,context:t,renderless:Y,api:K})}}),$e=["aria-expanded","aria-controls","aria-describedby"],Ce=["id","tabindex"],Ae={class:"tiny-collapse-item__title__right"},we=["aria-hidden","aria-labelledby","id"],ke={class:"tiny-collapse-item__content"};function Ie(e,r,t,n,a,i){var s=y("icon-chevron-right"),o=y("collapse-transition");return w(),k("div",{class:p(["tiny-collapse-item",{"is-active":e.state.isActive,"is-disabled":e.disabled}])},[d("div",{role:"tab","aria-expanded":e.state.isActive,"aria-controls":"tiny-collapse-content-"+e.state.id,"aria-describedby":"tiny-collapse-content-"+e.state.id},[d("div",{class:p(["tiny-collapse-item__header",{focusing:e.state.focusing,"is-active":e.state.isActive}]),role:"button",id:"tiny-collapse-head-"+e.state.id,tabindex:e.disabled?void 0:0,onKeyup:r[2]||(r[2]=B(m(function(){return e.handleEnterClick&&e.handleEnterClick.apply(e,arguments)},["stop"]),["space","enter"])),onFocus:r[3]||(r[3]=function(){return e.handleFocus&&e.handleFocus.apply(e,arguments)}),onBlur:r[4]||(r[4]=function(l){return e.state.focusing=!1}),onClick:r[5]||(r[5]=function(){return e.handleHeaderClick&&e.handleHeaderClick.apply(e,arguments)})},[d("div",{class:"tiny-collapse-item__arrow",onClick:r[0]||(r[0]=m(function(){return e.handleHeaderClick&&e.handleHeaderClick.apply(e,arguments)},["stop"]))},[c(e.$slots,"icon",{},function(){return[h(s,{class:p(["tiny-svg-size",{"is-active":e.state.isActive}])},null,8,["class"])]})]),d("div",{class:"tiny-collapse-item__word-overflow",onClick:r[1]||(r[1]=m(function(){return e.handleHeaderClick&&e.handleHeaderClick.apply(e,arguments)},["stop"]))},[c(e.$slots,"title",{},function(){return[g($(e.title),1)]})]),d("div",Ae,[c(e.$slots,"title-right",{},function(){return[g($(e.titleRight),1)]})])],42,Ce),h(o,null,{default:R(function(){return[P(d("div",{class:"tiny-collapse-item__wrap",role:"tabpanel","aria-hidden":!e.state.isActive,"aria-labelledby":"tiny-collapse-head-"+e.state.id,id:"tiny-collapse-content-"+e.state.id},[d("div",ke,[c(e.$slots,"default")])],8,we),[[L,e.state.isActive]])]}),_:3})],8,$e)],2)}var j=he(ge,[["render",Ie]]);function _e(e,r){var t=typeof Symbol<"u"&&e[Symbol.iterator]||e["@@iterator"];if(t)return(t=t.call(e)).next.bind(t);if(Array.isArray(e)||(t=Oe(e))||r&&e&&typeof e.length=="number"){t&&(e=t);var n=0;return function(){return n>=e.length?{done:!0}:{done:!1,value:e[n++]}}}throw new TypeError(`Invalid attempt to iterate non-iterable instance.
In order to be iterable, non-array objects must have a [Symbol.iterator]() method.`)}function Oe(e,r){if(e){if(typeof e=="string")return T(e,r);var t=Object.prototype.toString.call(e).slice(8,-1);if(t==="Object"&&e.constructor&&(t=e.constructor.name),t==="Map"||t==="Set")return Array.from(e);if(t==="Arguments"||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(t))return T(e,r)}}function T(e,r){(r==null||r>e.length)&&(r=e.length);for(var t=0,n=new Array(r);t<r;t++)n[t]=e[t];return n}var Ne=function(r,t){for(var n=r.__vccOpts||r,a=_e(t),i;!(i=a()).done;){var s=i.value,o=s[0],l=s[1];n[o]=l}return n},Ee=f({props:[].concat(C,["title","titleRight","name","disabled"]),components:{CollapseTransition:M,IconArrowBottom:J()},setup:function(r,t){return A({props:r,context:t,renderless:Y,api:K})}}),Se={"data-tag":"tiny-collapse-item",class:"border-b-0.5 sm:border-b border-solid border-color-border-disabled last:border-0"},je=["aria-expanded","aria-controls","aria-describedby"],Te=["id","tabindex"],He={"data-tag":"tiny-collapse-item-titleright",class:"text-xs font-normal"},Ve=["aria-hidden","aria-labelledby","id"],Fe={class:"pb-6 pt-0 px-4 sm:pt-0 sm:pr-0 sm:pl-4 sm:pb-4 text-xs text-color-text-primary leading-relaxed"};function Me(e,r,t,n,a,i){var s=y("icon-arrow-bottom"),o=y("collapse-transition");return w(),k("div",Se,[d("div",{role:"tab","data-tag":"tiny-collapse-item-tab","aria-expanded":e.state.isActive,"aria-controls":"tiny-collapse-content-"+e.state.id,"aria-describedby":"tiny-collapse-content-"+e.state.id},[d("div",{"data-tag":"tiny-collapse-item-body",class:p(["flex relative items-center h-12 sm:h-10 py-0 sm:pr-0 pl-4 pr-10 text-sm font-bold justify-between whitespace-nowrap",e.disabled?"text-color-text-disabled":"text-color-text-primary sm:[&:has(.peer:hover)_[role=title]]:text-color-brand"]),id:"tiny-collapse-head-"+e.state.id,tabindex:e.disabled?void 0:0,onKeyup:r[2]||(r[2]=B(m(function(){return e.handleEnterClick&&e.handleEnterClick.apply(e,arguments)},["stop"]),["space","enter"])),onFocus:r[3]||(r[3]=function(){return e.handleFocus&&e.handleFocus.apply(e,arguments)}),onBlur:r[4]||(r[4]=function(l){return e.state.focusing=!1})},[d("div",{"data-tag":"tiny-collapse-item-title",class:p(["whitespace-nowrap overflow-hidden overflow-ellipsis inline-block peer",[e.disabled?"cursor-not-allowed":"cursor-pointer sm:hover:text-color-brand"]]),role:"title",onClick:r[0]||(r[0]=function(){return e.handleHeaderClick&&e.handleHeaderClick.apply(e,arguments)})},[c(e.$slots,"title",{},function(){return[g($(e.title),1)]})],2),d("div",He,[c(e.$slots,"title-right",{},function(){return[g($(e.titleRight),1)]})]),d("div",{"data-tag":"tiny-collapse-item-icon",class:p(["absolute sm:left-0 right-3.5 text-xs mr-1 w-3 peer",[e.disabled?"fill-color-text-disabled  cursor-not-allowed":"fill-color-icon-secondary cursor-pointer sm:peer-hover:fill-color-brand sm:hover:fill-color-brand"]]),onClick:r[1]||(r[1]=function(){return e.handleHeaderClick&&e.handleHeaderClick.apply(e,arguments)})},[c(e.$slots,"icon",{active:e.state.isActive,disabled:e.disabled},function(){return[h(s,{"custom-class":"w-2.5 h-2.5 sm:w-3 sm:h-3 transition-transform duration-300",class:p([e.state.isActive?"sm:rotate-0 rotate-180":"sm:-rotate-90 rotate-0"])},null,8,["class"])]})],2)],42,Te)],8,je),h(o,null,{default:R(function(){return[P(d("div",{"data-tag":"tiny-collapse-item-active",class:"will-change-[height] bg-color-bg-1 overflow-hidden box-border",role:"tabpanel","aria-hidden":!e.state.isActive,"aria-labelledby":"tiny-collapse-head-"+e.state.id,id:"tiny-collapse-content-"+e.state.id},[d("div",Fe,[c(e.$slots,"default")])],8,Ve),[[L,e.state.isActive]])]}),_:3})])}var Be=Ne(Ee,[["render",Me]]);function _(){return _=Object.assign?Object.assign.bind():function(e){for(var r=1;r<arguments.length;r++){var t=arguments[r];for(var n in t)Object.prototype.hasOwnProperty.call(t,n)&&(e[n]=t[n])}return e},_.apply(this,arguments)}var Re=function(r){var t,n=typeof process=="object"?(t={})==null?void 0:t.TINY_MODE:null;return(n||r)==="pc"?j:(n||r)==="mobile-first"?Be:j},Pe={INTERVAL:50},Le=_({},H,{_constants:{type:Object,default:function(){return Pe}},title:String,titleRight:String,name:{type:[String,Number],default:function(){return this._uid}},disabled:Boolean}),b=f({name:V+"CollapseItem",componentName:"CollapseItem",inject:["collapse"],props:Le,setup:function(r,t){return F({props:r,context:t,template:Re})}}),De="3.15.0";b.install=function(e){e.component(b.name,b)};b.version=De;export{b as C,v as a};
