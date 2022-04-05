<template>

<el-container>
  <el-header class="header">
      <h3>Programming Graph</h3>
      <el-button type="primary"   @click="exportEditor">Export</el-button>
  </el-header>
  <el-container class="container">
    <el-aside width="200px" class="column">
        <ul>
            <li v-for="n in listNodes" :key="n" draggable="true" :data-node="n.item" @dragstart="drag($event)" class="drag-drawflow" >
                <div class="node" :style="`background: ${n.color}`" >{{ n.name }}</div>
            </li>
        </ul>
    </el-aside>
    <el-main>
        <div id="drawflow" @drop="drop($event)" @dragover="allowDrop($event)"></div>
    </el-main>
  </el-container>
</el-container>
<el-dialog
    v-model="dialogVisible"
    title="Export"
    width="50%"
  >
    <span>Data:</span>
    <pre><code>{{dialogData}}</code></pre>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisible = false">Cancel</el-button>
        <el-button type="primary" @click="dialogVisible = false"
          >Confirm</el-button
        >
      </span>
    </template>
  </el-dialog>
</template>
<script>

import Drawflow from 'drawflow'
import styleDrawflow from 'drawflow/dist/drawflow.min.css'
import style from '../assets/style.css' 
import { onMounted, shallowRef, h, getCurrentInstance, render, readonly, ref } from 'vue'
import NodeNumber from './nodes/nodeNumber.vue'
import NodeBoolean from './nodes/nodeBoolean.vue'
import NodeVariable from './nodes/nodeVariable.vue'
import NodeArithmetic from './nodes/nodeArithmetic.vue'
import NodeComparison from './nodes/nodeComparison.vue'
import NodeLogical from './nodes/nodeLogical.vue'

export default {
  name: 'drawflow',
  setup() {
   const listNodes = readonly([
        {
            name: 'Boolean',
            color: '#4c4c4c',
            item: 'NodeBoolean',
            input:0,
            output:1
        },
        {
            name: 'Number',
            color: '#4c4c4c',
            item: 'NodeNumber',
            input:0,
            output:1
        },
        {
            name: 'Variable',
            color: '#4c4c4c',
            item: 'NodeVariable',
            input:0,
            output:1
        },
        {
            name: 'Arithmetic Operator',
            color: '#4c4c4c',
            item: 'NodeArithmetic',
            input:2,
            output:1
        },
        {
            name: 'Comparison Operator',
            color: '#4c4c4c',
            item: 'NodeComparison',
            input:2,
            output:1
        },
        {
            name: 'Logical Operator',
            color: '#4c4c4c',
            item: 'NodeLogical',
            input:2,
            output:1
        },
        {
            name: 'Assign',
            color: '#4c4c4c',
            item: 'NodeBoolean',
            input:2,
            output:1
        },
        {
            name: 'Print',
            color: '#4c4c4c',
            item: 'NodeBoolean',
            input:1,
            output:1
        },
        {
            name: 'Conditional',
            color: '#4c4c4c',
            item: 'NodeBoolean',
            input:1,
            output:3
        },
        {
            name: 'Loop',
            color: '#4c4c4c',
            item: 'NodeBoolean',
            input:3,
            output:2
        },
    ])
   
   const editor = shallowRef({})
   const dialogVisible = ref(false)
   const dialogData = ref({})
   const Vue = { version: 3, h, render };
   const internalInstance = getCurrentInstance()
   internalInstance.appContext.app._context.config.globalProperties.$df = editor;
   
    function exportEditor() {
      dialogData.value = editor.value.export();
      dialogVisible.value = true;
    }

    const drag = (ev) => {
      if (ev.type === "touchstart") {
        mobile_item_selec = ev.target.closest(".drag-drawflow").getAttribute('data-node');
      } else {
      ev.dataTransfer.setData("node", ev.target.getAttribute('data-node'));
      }
    }
    const drop = (ev) => {
      if (ev.type === "touchend") {
        var parentdrawflow = document.elementFromPoint( mobile_last_move.touches[0].clientX, mobile_last_move.touches[0].clientY).closest("#drawflow");
        if(parentdrawflow != null) {
          addNodeToDrawFlow(mobile_item_selec, mobile_last_move.touches[0].clientX, mobile_last_move.touches[0].clientY);
        }
        mobile_item_selec = '';
      } else {
        ev.preventDefault();
        var data = ev.dataTransfer.getData("node");
        addNodeToDrawFlow(data, ev.clientX, ev.clientY);
      }

    }
    const allowDrop = (ev) => {
      ev.preventDefault();
    }

   let mobile_item_selec = '';
   let mobile_last_move = null;
   function positionMobile(ev) {
     mobile_last_move = ev;
   }

    function addNodeToDrawFlow(name, pos_x, pos_y) {
      pos_x = pos_x * ( editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().x * ( editor.value.precanvas.clientWidth / (editor.value.precanvas.clientWidth * editor.value.zoom)));
      pos_y = pos_y * ( editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)) - (editor.value.precanvas.getBoundingClientRect().y * ( editor.value.precanvas.clientHeight / (editor.value.precanvas.clientHeight * editor.value.zoom)));
    
      const nodeSelected = listNodes.find(ele => ele.item == name);
      editor.value.addNode(name, nodeSelected.input,  nodeSelected.output, pos_x, pos_y, name, {}, name, 'vue');
      
    }


   onMounted(() => {

      var elements = document.getElementsByClassName('drag-drawflow');
      for (var i = 0; i < elements.length; i++) {
        elements[i].addEventListener('touchend', drop, false);
        elements[i].addEventListener('touchmove', positionMobile, false);
        elements[i].addEventListener('touchstart', drag, false );
      }
        
       const id = document.getElementById("drawflow");
       editor.value = new Drawflow(id, Vue, internalInstance.appContext.app._context);
       editor.value.start();
       
       editor.value.registerNode('NodeNumber', NodeNumber, {}, {});
       editor.value.registerNode('NodeBoolean', NodeBoolean, {}, {});
       editor.value.registerNode('NodeVariable', NodeVariable, {}, {});
       editor.value.registerNode('NodeArithmetic', NodeArithmetic, {}, {});
       editor.value.registerNode('NodeComparison', NodeComparison, {}, {});
       editor.value.registerNode('NodeLogical', NodeLogical, {}, {});

       editor.value.on('nodeDataChanged', (id) => {
			let checkedNode = editor.value.getNodeFromId(id);

			let count = 0;
			let lastKey;
			for(var key in checkedNode.inputs) {
				lastKey = key;
				if(checkedNode.inputs.hasOwnProperty(key)) {
					count++;
				}
			}

			if (checkedNode.name == 'NodeArithmetic' || checkedNode.name == 'NodeLogical') {
				if (checkedNode.data.data.element == '- ' || checkedNode.data.data.element == ' NOT ') {
					if (count > 1) {
						editor.value.removeNodeInput(id, lastKey);
					}
				} else {
					if (count < 2) {
						editor.value.addNodeInput(id);
					}
				}
			}
       });

       editor.value.on('connectionCreated', (link) => {
			let output = editor.value.getNodeFromId(link.output_id);
			let input = editor.value.getNodeFromId(link.input_id);
			if ( input.inputs[link.input_class].connections.length > 1 || output.outputs[link.output_class].connections.length > 1 ||
				(input.name == 'NodeLogical' && output.name != 'NodeVariable' && output.name != 'NodeBoolean' 
					&& output.name != 'NodeComparison' && output.name != 'NodeLogical') ||
				(input.name == 'NodeComparison' && output.name != 'NodeVariable' && output.name != 'NodeNumber' && output.name != 'NodeArithmetic') ||
				(input.name == 'NodeArithmetic' && output.name != 'NodeVariable' && output.name != 'NodeNumber' && output.name != 'NodeArithmetic')) {
				editor.value.removeSingleConnection(link.output_id, link.input_id, link.output_class, link.input_class)
			}
       });

       editor.value.import({"drawflow":{"Home":{"data":{}}}})})

  return {
    exportEditor, listNodes, drag, drop, allowDrop, dialogVisible, dialogData
  }

  }
}

</script>
<style scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    border-bottom: 1px solid #494949;
}
.container {
    min-height: calc(100vh - 100px);
}
.column {
    border-right: 1px solid #494949;
}
.column ul {
    padding-inline-start: 0px;
    padding: 10px 10px;
 
}
.column li {
    background: transparent;
}

.node {
    border-radius: 8px;
    border: 2px solid #494949;
    display: block;
    height:60px;
    line-height:40px;
    padding: 10px;
    margin: 10px 0px;
    cursor: move;

}
#drawflow {
  width: 100%;
  height: 100%;
  text-align: initial;
  background: #2b2c30;
  background-size: 20px 20px;
  background-image: radial-gradient(#494949 1px, transparent 1px);
  
}
</style>
