export const getActionColor = (action: string) => {
	const colorMap = {
		CREATED: "blue",
		ACCEPT: "green",
		FORWARD: "orange",
		RETURN: "yellow",
		TRANSFER: "purple",
		COMPLETE: "positive",
		REJECTED: "negative",
		TO_APPROVE: "info",
		APPROVED: "green",
	};
	return colorMap[action as keyof typeof colorMap] || "grey";
};

export const getStatusColor = (status: string) => {
	const colorMap = {
		COMPLETED: "green",
		IN_PROGRESS: "blue",
		UNDER_REVIEW: "orange",
		REJECTED: "red",
	};
	return colorMap[status as keyof typeof colorMap] || "grey";
};
