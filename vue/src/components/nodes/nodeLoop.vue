<template>
    <div ref="el">
        <nodeHeader title="Loop"/>
		<el-tag size="small">Instruction</el-tag>
		<br>
		<el-tag size="small" class="tag_float">External Instruction</el-tag>
		<br>
		<el-tag size="small">Variable</el-tag>
		<br>
		<el-tag size="small" class="tag_float">Internal Instruction</el-tag>
		<br>
		<el-tag size="small">Number</el-tag>
		<br><br>
		<el-tag size="small">Number</el-tag>
		<br><br>
    </div>
</template>

<script>
import { defineComponent, onMounted, getCurrentInstance, ref, nextTick } from 'vue'
import nodeHeader from './nodeHeader.vue'

export default defineComponent({
    components: {
        nodeHeader
    },
    setup() {
        const el = ref(null);
        const nodeId = ref(0);
        let df = null
        const condition = ref('for');
        const dataNode = ref({});
        
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
			dataNode.value.data.element = 'for';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el, condition, updateSelect
        }

    }    
    
})
</script>
