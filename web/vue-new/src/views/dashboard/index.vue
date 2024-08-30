


<template>
	<div>
	  <el-row>
		<el-col :span="24">
		  <h1>欢迎光临!</h1>
		</el-col>
	  </el-row>
  
	  <el-row>
		<el-col :span="24">
		  <el-carousel height="540px">
			<el-carousel-item>
				<img src="../../assets/swiper_01.jpg">
			</el-carousel-item>
			<el-carousel-item>
				<img src="../../assets/swiper_02.jpg">
			</el-carousel-item>
			<el-carousel-item>
				<img src="../../assets/swiper_03.jpg">
			</el-carousel-item>
		  </el-carousel>
		</el-col>
	  </el-row>
	</div>
  </template>
  
  <script>
  import { Carousel, CarouselItem, Row, Col } from 'element-ui';
  
  export default {
	name: 'HomePage',
	components: {
	  'el-carousel': Carousel,
	  'el-carousel-item': CarouselItem,
	  'el-row': Row,
	  'el-col': Col
	}
  }
  </script>
  
  <style scoped>
  .el-carousel {
	margin-top: 15px;
  }
  
  .el-carousel-item img {
	width: 100%;
	height: 100%;
	object-fit: cover;
  }
  </style>
  
