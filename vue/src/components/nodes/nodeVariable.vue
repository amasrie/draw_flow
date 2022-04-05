<template>
    <div ref="el">
        <nodeHeader  title="Variable"/>
        <el-input v-model="variable" minlength="1" placeholder="Variable name" @change="updateSelect" size="small" df-variable></el-input>
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
        const variable = ref("");
        const dataNode = ref({});
		const regExp = /^[a-zA-Z_][a-zA-Z_0-9]*$/;
		let hasError = true;
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
			el.value.children[1].children[0].classList.value = el.value.children[1].children[0].classList.value.replace(' inputError', ''); 
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value);
			hasError = !regExp.test(variable.value);
			el.value.children[1].children[0].classList.value += hasError ? ' inputError' : '';
        }

        onMounted(async () => {
            await nextTick()
			el.value.children[1].children[0].classList.value += ' inputError';
            nodeId.value = el.value.parentElement.parentElement.id.slice(5)
            dataNode.value = df.getNodeFromId(nodeId.value)
            variable.value = dataNode.value.data.element;
        });

        return {
            el, variable, updateSelect
        }

    }    
    
})
</script>
