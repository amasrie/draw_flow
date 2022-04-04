<template>
    <div ref="el">
        <nodeHeader  title="Number"/>
        <el-input-number v-model="numbers" placeholder="Numeric value" @change="updateSelect" size="small" df-numbers></el-input-number>
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
        const numbers = ref('get');
		numbers.value = 0;
        const dataNode = ref({});
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value);
			if (!numbers.value) {
				numbers.value = 0;
			}
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
            
            numbers.value = dataNode.value.data.numbers;
			if (!numbers.value) {
				numbers.value = 0;
			}
        });

        return {
            el, numbers, updateSelect
        }

    }    
    
})
</script>
