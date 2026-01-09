import { ref, onMounted, onUnmounted, readonly } from 'vue';

export function useBreakpoints() {
	const isMobile = ref(typeof window !== 'undefined' ? window.innerWidth < 1100 : false);

	const updateBreakpoint = () => {
		isMobile.value = window.innerWidth < 1100;
	};

	onMounted(() => {
		window.addEventListener('resize', updateBreakpoint);
	});

	onUnmounted(() => {
		window.removeEventListener('resize', updateBreakpoint);
	});

	return { 
		isMobile: readonly(isMobile) 
	};
}