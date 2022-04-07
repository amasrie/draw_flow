<template>
    <div ref="el">
        <nodeHeader title="Conditional"/>
		<br>
		<el-tag size="small">Instruction</el-tag>
		<el-tag size="small" class="tag_float">External Instruction</el-tag>
		<br><br>
		<el-tag v-show="isBinary" size="small">Boolean</el-tag>
		<el-tag size="small" class="tag_float">Internal Instruction</el-tag>
		<br><br>
        <el-select v-model="condition" placeholder="Select condition" @change="updateSelect" size="small" df-condition>
		    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"></el-option>
	    </el-select>
    </div>
</template>

<script>
import { defineComponent, onMounted, getCurrentInstance, ref, nextTick, readonly } from 'vue'
import nodeHeader from './nodeHeader.vue'

export default defineComponent({
    components: {
        nodeHeader
    },
    setup() {
        const el = ref(null);
        const nodeId = ref(0);
        let df = null
        const condition = ref('if');
        const isBinary = ref(true);
        const dataNode = ref({});
        const options = readonly([
            {
                value: 'if',
                label: 'First condition'
            },
            {
                value: 'elif',
                label: 'Next condition'
            },
            {
                value: 'else',
                label: 'Otherwise'
            },
        ]);
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
			df.dispatch('nodeDataChanged', nodeId.value );
			condition.value = value;
			isBinary.value = value =='else' ? false : true;
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = 'if';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el, condition, updateSelect, options, isBinary
        }

    }    
    
})
</script>
