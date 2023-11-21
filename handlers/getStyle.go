package handlers

func getStyle() string {
	return `
<defs>
	<style>
		.title{
			font-family: 'Microsoft YaHei Black', sans-serif;
    		font-size : 20px;
    		text-anchor : middle;
    		fill : black;
    		font-weight : bold;
			transform: translateY(0.4em);
		}
		.header{
   			fill : white;
    		stroke : black;
    		stroke-width : 1;
    		stroke-opacity : 1.0;
		}
		.header-text{
			font-family: 'Microsoft YaHei Black', sans-serif;
    		font-size : 16px;
    		text-anchor : middle;
    		fill : black;
    		font-weight : normal;
			transform: translateY(0.4em);
		}
		.cell{
   			fill : white;
    		stroke : black;
    		stroke-width : 1;
    		stroke-opacity : 1.0;
		}

		/* Light Salmon */
		.color-fill-0 {
			fill: #FFA07A;
		}
		
		/* Coral */
		.color-fill-1 {
			fill: #FF7F50;
		}
		
		/* Tomato */
		.color-fill-2 {
			fill: #FF6347;
		}
		
		/* Orange */
		.color-fill-3 {
			fill: #FFA500;
		}
		
		/* Gold */
		.color-fill-4 {
			fill: #FFD700;
		}
		
		/* Yellow */
		.color-fill-5 {
			fill: #FFFF00;
		}
		
		/* Green Yellow */
		.color-fill-6 {
			fill: #ADFF2F;
		}
		
		/* Lime Green */
		.color-fill-7 {
			fill: #32CD32;
		}
		
		/* Turquoise */
		.color-fill-8 {
			fill: #40E0D0;
		}
		
		/* Sky Blue */
		.color-fill-9 {
			fill: #87CEEB;
		}
		
		/* Dodger Blue */
		.color-fill-10 {
			fill: #1E90FF;
		}
		
		/* Medium Slate Blue */
		.color-fill-11 {
			fill: #7B68EE;
		}
		
		/* Blue Violet */
		.color-fill-12 {
			fill: #8A2BE2;
		}
		
		/* Medium Orchid */
		.color-fill-13 {
			fill: #BA55D3;
		}
		
		/* Hot Pink */
		.color-fill-14 {
			fill: #FF69B4;
		}
		
		/* Deep Pink */
		.color-fill-15 {
			fill: #FF1493;
		}
	</style>
</defs>
`
}
