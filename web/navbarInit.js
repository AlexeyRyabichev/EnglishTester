// document.addEventListener('DOMContentLoaded', function () {
//     let elements = document.querySelectorAll('.sidenav');
//     let instances = M.Sidenav.init(elements, {});
//     import docx4js from "docx4js"
//
//     docx4js.load("~/test.docx").then(docx=>{
//         //you can render docx to anything (react elements, tree, dom, and etc) by giving a function
//         docx.render(function createElement(type,props,children){
//             return {type,props,children}
//         });
//
//         //or use a event handler for more flexible control
//         const ModelHandler=require("docx4js/openxml/docx/model-handler").default;
//         class MyModelhandler extends ModelHandler{
//             onp({type,children,node,...}, node, officeDocument){
//
//             }
//         }
//         let handler=new MyModelhandler();
//         handler.on("*",function({type,children,node,...}, node, officeDocument){
//             console.log("found model:"+type)
//         });
//         handler.on("r",function({type,children,node,...}, node, officeDocument){
//             console.log("found a run")
//         });
//
//         docx.parse(handler);
//
//         //you can change content on docx.officeDocument.content, and then save
//         // docx.officeDocument.content("w\\:t").text("hello");
//         // docx.save("~/changed.docx")
//     })
// });