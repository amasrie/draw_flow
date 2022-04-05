<template>
    <div ref="el">
        <nodeHeader  title="Arithmetic Operator"/>
        <el-select v-model="operator" placeholder="Select operator" @change="updateSelect" size="small" df-operator>
		    <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value"></el-option>
	    </el-select>

    <br><br>
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
        const operator = ref(' + ');
        const dataNode = ref({});
        const options = readonly([
            {
                value: ' + ',
                label: 'Add'
            },
            {
                value: '- ',
                label: 'Negative'
            },
            {
                value: ' - ',
                label: 'Subtract'
            },
            {
                value: ' * ',
                label: 'Multiply'
            },
            {
                value: ' / ',
                label: 'Divide'
            },
            {
                value: ' // ',
                label: 'Divide (Rounded)'
            },
            {
                value: ' % ',
                label: 'Modulo'
            },
            {
                value: ' ** ',
                label: 'Exponent'
            },
        ]);
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value);
			df.dispatch('nodeDataChanged', nodeId.value );
			operator.value = value;
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = ' -';
            df.updateNodeDataFromId(nodeId.value, dataNode.value);
        });

        return {
            el, operator, updateSelect, options
        }

    }    
    
})
</script>
