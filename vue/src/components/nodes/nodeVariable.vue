<template>
    <div ref="el">
        <nodeHeader title="Variable"/>
		<el-tag size="small" class="tag_float">Anything</el-tag>
		<br><br>
        <el-input v-model="variable" minlength="1" placeholder="Variable name" @change="updateSelect" size="small" df-variable></el-input>
    	<br>
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
		const reserved = [
			'False', 'None', 'True', '__peg_parser__', 'and', 'assert', 'as', 'async', 'await',
			'break', 'class', 'continue', 'def', 'del', 'elif', 'else', 'except', 'finally',
			'for', 'from', 'global', 'if', 'import', 'in', 'is', 'lambda', 'nonlocal',
			'not', 'or', 'pass', 'raise', 'return', 'try', 'while', 'with', 'yield', 'print', 'exec'
		]
		let hasError = true;
        
        df = getCurrentInstance().appContext.config.globalProperties.$df.value;

        const updateSelect = (value) => {
			el.value.children[4].children[0].classList.value = el.value.children[4].children[0].classList.value.replace(' inputError', ''); 
            dataNode.value.data.element = value;
            df.updateNodeDataFromId(nodeId.value, dataNode.value.data);
			hasError = !regExp.test(variable.value) || reserved.includes(variable.value);
			el.value.children[4].children[0].classList.value += hasError ? ' inputError' : '';
        }

        onMounted(async () => {
            await nextTick()
			el.value.children[4].children[0].classList.value += ' inputError';
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
