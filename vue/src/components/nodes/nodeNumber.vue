<template>
    <div ref="el">
        <nodeHeader  title="Number"/>
        <el-input-number v-model="numbers" :precision="5" :step="0.1" placeholder="Numeric value" @change="updateSelect" size="small" df-numbers></el-input-number>
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
        const numbers = ref(0);
        const dataNode = ref({});
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
			numbers.value = value;
			if (!numbers.value) {
				numbers.value = 0;
			}
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
            
            numbers.value = dataNode.value.data.element;
			if (!numbers.value) {
				numbers.value = 0;
				dataNode.value.data.element = 0;
	            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
			}
        });

        return {
            el, numbers, updateSelect
        }

    }    
    
})
</script>
