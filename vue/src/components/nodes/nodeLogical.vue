<template>
    <div ref="el">
        <nodeHeader title="Logical Operator"/>
		<el-tag size="small">Boolean</el-tag>
		<br v-show="isBinary">
		<el-tag size="small" class="tag_float">Boolean</el-tag>
		<br v-show="isBinary">
		<el-tag v-show="isBinary" size="small">Boolean</el-tag>
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
        const operator = ref(' and ');
        const isBinary = ref(true);
        const dataNode = ref({});
        const options = readonly([
            {
                value: ' and ',
                label: 'Conjunction'
            },
            {
                value: ' or ',
                label: 'Disjunction'
            },
            {
                value: ' not ',
                label: 'Negation'
            },
        ]);
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
			df.dispatch('nodeDataChanged', nodeId.value );
			operator.value = value;
			isBinary.value = value ==' not ' ? false : true;
        }

        onMounted(async () => {
            await nextTick()
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
			dataNode.value.data.element = ' and ';
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
        });

        return {
            el, operator, updateSelect, options, isBinary
        }

    }    
    
})
</script>
