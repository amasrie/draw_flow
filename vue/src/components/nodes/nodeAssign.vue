<template>
    <div ref="el">
        <nodeHeader title="Assign"/>
		<el-tag size="small">Instruction</el-tag>
		<br><br>
		<el-tag size="small">Variable</el-tag>
		<el-tag size="small" class="tag_float">Instruction</el-tag>
		<br><br>
		<el-tag size="small">Any Expression</el-tag>
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
        const dataNode = ref({});
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = ' = ';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el
        }

    }    
    
})
</script>
