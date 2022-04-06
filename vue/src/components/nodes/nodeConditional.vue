<template>
    <div ref="el">
        <nodeHeader  title="Conditional"/>
        <el-select v-model="condition" placeholder="Select condition" @change="updateSelect" size="small" df-condition>
		    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"></el-option>
	    </el-select>

    <br><br><br><br>
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
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = 'if';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el, condition, updateSelect, options
        }

    }    
    
})
</script>
