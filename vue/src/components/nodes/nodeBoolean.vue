<template>
    <div ref="el">
        <nodeHeader  title="Boolean"/>
        <el-switch active-text="True" inactive-text="False" v-model="boolean" @change="updateSelect" df-boolean></el-switch>
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
        const boolean = ref('False');
        const dataNode = ref({});
		let first = true;
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = async (value) => {
			if (!first) {
		        dataNode.value.data.element = value ? "True" : "False";
		        df.updateNodeDataFromId(nodeId.value, dataNode.value);
			}
			first = false;
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
            
            boolean.value = dataNode.value.data.element;
			if (!boolean.value) {
				boolean.value = 'False';
				dataNode.value.data.element = 'False';
	            df.updateNodeDataFromId(nodeId.value, dataNode.value);
			}
        });

        return {
            el, boolean, updateSelect
        }

    }    
    
})
</script>
