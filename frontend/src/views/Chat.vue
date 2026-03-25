<script setup>
import { ref, onMounted } from 'vue'
import { useUserStore } from '@/stores/user'
import { qaAPI } from '@/api'

const userStore = useUserStore()

const messages = ref([])
const inputMessage = ref('')
const loading = ref(false)
const conversationId = ref('')

const sendMessage = async () => {
  if (!inputMessage.value.trim() || loading.value) return
  
  const userMsg = { role: 'user', content: inputMessage.value }
  messages.value.push(userMsg)
  const question = inputMessage.value
  inputMessage.value = ''
  
  loading.value = true
  try {
    const userContext = userStore.profile ? {
      score: userStore.profile.score,
      rank: userStore.profile.rank,
      preferred_subjects: userStore.profile.preferred_subjects,
      preferred_regions: userStore.profile.preferred_regions,
      interest_tags: userStore.profile.interest_tags
    } : null
    
    const res = await qaAPI.chat({
      message: question,
      conversation_id: conversationId.value,
      user_context: userContext
    })
    
    conversationId.value = res.data.conversation_id
    
    const assistantMsg = {
      role: 'assistant',
      content: res.data.answer,
      sources: res.data.sources,
      recommendations: res.data.recommendations
    }
    messages.value.push(assistantMsg)
  } catch (error) {
    messages.value.push({
      role: 'assistant',
      content: '抱歉，发生了错误，请稍后再试。'
    })
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  messages.value.push({
    role: 'assistant',
    content: '您好！我是高考志愿填报指导助手。请告诉我您的分数、位次、偏好专业等信息，我可以为您推荐合适的院校和专业，提供志愿填报建议。'
  })
})
</script>

<template>
  <div class="chat-container">
    <el-card class="chat-card">
      <template #header>
        <h2>智能问答</h2>
      </template>
      
      <div class="chat-messages" v-loading="loading">
        <div
          v-for="(msg, index) in messages"
          :key="index"
          :class="['message', msg.role]"
        >
          <div class="message-avatar">
            {{ msg.role === 'user' ? '我' : 'AI' }}
          </div>
          <div class="message-content">
            <div class="content-text">{{ msg.content }}</div>
            <div v-if="msg.recommendations" class="recommendations">
              <el-divider content-position="left">推荐</el-divider>
              <div v-if="msg.recommendations.colleges" class="rec-section">
                <h4>推荐院校：</h4>
                <div v-for="c in msg.recommendations.colleges" :key="c.name" class="rec-item">
                  {{ c.name }} - {{ c.probability }}%录取概率 - {{ c.reason }}
                </div>
              </div>
              <div v-if="msg.recommendations.majors" class="rec-section">
                <h4>推荐专业：</h4>
                <div v-for="m in msg.recommendations.majors" :key="m.name" class="rec-item">
                  {{ m.name }} - 匹配度{{ m.match_score }}%
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      
      <div class="chat-input">
        <el-input
          v-model="inputMessage"
          placeholder="请输入您的问题..."
          @keyup.enter="sendMessage"
        >
          <template #append>
            <el-button @click="sendMessage" :disabled="loading">发送</el-button>
          </template>
        </el-input>
      </div>
    </el-card>
  </div>
</template>

<style scoped>
.chat-container {
  max-width: 800px;
  margin: 0 auto;
}

.chat-card {
  height: calc(100vh - 200px);
  display: flex;
  flex-direction: column;
}

.chat-messages {
  flex: 1;
  overflow-y: auto;
  padding: 20px;
}

.message {
  display: flex;
  margin-bottom: 20px;
}

.message.user {
  flex-direction: row-reverse;
}

.message-avatar {
  width: 40px;
  height: 40px;
  border-radius: 50%;
  background: #409eff;
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.message.assistant .message-avatar {
  background: #67c23a;
}

.message-content {
  max-width: 70%;
  margin: 0 10px;
}

.content-text {
  padding: 10px 15px;
  border-radius: 10px;
  line-height: 1.6;
}

.message.user .content-text {
  background: #409eff;
  color: white;
}

.message.assistant .content-text {
  background: #f5f5f5;
  color: #333;
}

.recommendations {
  margin-top: 10px;
}

.rec-section {
  margin-bottom: 10px;
}

.rec-section h4 {
  margin: 5px 0;
}

.rec-item {
  padding: 5px 10px;
  background: #f0f9eb;
  border-radius: 5px;
  margin: 5px 0;
  font-size: 14px;
}

.chat-input {
  padding: 10px;
  border-top: 1px solid #eee;
}
</style>
