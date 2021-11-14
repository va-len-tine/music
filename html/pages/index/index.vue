<template>
	<view class="container">
		<jmSearch v-on:flush_data="flush_data"></jmSearch>
		<view class="">
			<radio-group @change="switch_search">
				<label>
					<radio checked="true" value="migu" /><text>咪咕音乐</text>
				</label>
				<label>
					<radio value="qq" /><text>QQ音乐</text>
				</label>
				<label>
					<radio value="netease" /><text>网易云音乐</text>
				</label>
			</radio-group>
		</view>
		<view class="">
			<button class="bt" size="mini" type="default" @click="get_list('19723756')">飙升榜</button>
			<button class="bt" size="mini" type="default" @click="get_list('3779629')">新歌榜</button>
			<button class="bt" size="mini" type="default" @click="get_list('2884035')">原创榜</button>
			<button class="bt" size="mini" type="default" @click="get_list('3778678')">热歌榜</button>
			<button class="bt" size="mini" type="default" @click="confirmDialog"><text class="button-text">链接</text></button>
		</view>
		<uni-popup id="dialogInput" ref="dialogInput" type="dialog">
			<uni-popup-dialog mode="input" title="下载链接" :value="audio[now]" placeholder="请输入内容" @confirm="dialogInputConfirm"></uni-popup-dialog>
		</uni-popup>
		<scroll-view scroll-top="0" scroll-y="true" class="scroll-Y">
			<uni-list v-for="(item, index) in dataList" :key=index>
				<uni-list-item :class = "index == now ? 'addclass' : '' " :title="item.name" thumb-size="medium" :thumb="item.pic" :rightText="item.author" clickable  @click="playone(index)"></uni-list-item>
			</uni-list>
	    </scroll-view>
		<imt-audio class="footer" continue :src="audio[now]" :duration="audio[now].duration" @prev="now = now === 0?audio.length-1:now-1" @next="now = now === audio.length-1?0:now+1"></imt-audio>
	</view>
</template>

<script>
	import jmSearch from "../../components/jm-search/jm-search.vue"
	import imtAudio from 'components/imt-audio/imt-audio'
	export default {
		components: {jmSearch,imtAudio},
		
		data() {
			return {
				dataList: [],			//播放列表
				audio: ['http://39.101.203.25/file/危险派对.mp3'],
				now: 0,					//默认播放歌曲
				search_type: "migu",	//默认搜索类型:migu
			}
		},
		
		methods: {
			//获取链接
			confirmDialog() {
				this.$refs.dialogInput.open()
			},
			
			//下载至服务器
			dialogInputConfirm(done, val) {
				this.value = val
				setTimeout(() => {
					uni.hideLoading()
					// uni.request({
					// 	url:'http://39.101.203.25/download/api',
					// 	method:'POST',
					// 	data:this.dataList[this.now],
					// 	success: (res) => {
					// 		console.log(res.statusCode)
					// 	}
					// });
					done()
				}, 0)
			},
			
			//将search_type传入搜索子组件
			switch_search(evt){
				this.search_type = evt.detail["value"]
			},        
			
			//搜索子组件返回歌曲，进行页面渲染
			flush_data(r){
				uni.request({
					url:'http://39.101.203.25/' + this.search_type + '/api?keyword=' + r,
					method:'GET',
					data:{},
					success: (res) => {
						this.now = 0
						this.dataList = res.data
						this.audio = res.data.map(x => {return x.src})
					}
				});
			},
			
			//播放选中歌曲
			playone(msg){
				this.now = msg
			},
			
			//飙升榜...
			get_list(s){
				uni.request({
					url:'http://39.101.203.25/netease/api?keyword=' + s,
					method:'GET',
					data:{},
					success: (res) => {
						this.dataList = res.data
						this.audio = res.data.map(x => {return x.src})
					}
				});
			},
		},
		
		//打开页面时，初始化数据
		beforeCreate() {
			uni.request({
				url:'http://39.101.203.25/netease/api?keyword=' + "抖音",
				method:'GET',
				data:{},
				success: (res) => {
					this.dataList = res.data
					this.audio = res.data.map(x => {return x.src})
				}
			});
		}
	}
</script>

<style>
	.scroll-Y {
			height: 850rpx;
		}
		
	.bt{
		background-color: #ffe3d9;
		color: #d91604;
		text-align: center;
		padding: 0.1% 0 0;
		width: 20%;
	}
	
	.addclass{
		background-color: #ffe3d9;
	}
	
	.footer{
		position: absolute;
		bottom: 0;
		width: 100%;
	}
</style>