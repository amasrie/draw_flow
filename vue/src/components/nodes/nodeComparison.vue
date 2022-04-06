<template>
    <div ref="el">
        <nodeHeader title="Comparison Operator"/>
		<el-tag size="small">Number</el-tag>
		<br>
		<el-tag size="small" class="tag_float">Boolean</el-tag>
		<br>
		<el-tag size="small">Number</el-tag>
		<br><br>
        <el-select v-model="operator" placeholder="Select operator" @change="updateSelect" size="small" df-operator>
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
        const operator = ref(' == ');
        const dataNode = ref({});
        const options = readonly([
            {
                value: ' == ',
                label: 'Equals'
            },
            {
                value: ' != ',
                label: 'Not equals'
            },
            {
                value: ' < ',
                label: 'Less than'
            },
            {
                value: ' <= ',
                label: 'Less than or equal'
            },
            {
                value: ' > ',
                label: 'Greater than'
            },
            {
                value: ' >= ',
                label: 'Greater than or equal'
            },
        ]);
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
			df.dispatch('nodeDataChanged', nodeId.value );
			operator.value = value;
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = ' == ';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el, operator, updateSelect, options
        }

    }    
    
})
</script>
