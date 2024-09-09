import axios from 'axios';

const API_URL = 'https://192.168.110.132:8000'; // 替换为你的 API URL

export default {
  async verifyToken(token: string): Promise<boolean> {
    try {
      const response = await axios.post(`${API_URL}/v1/info'`, { token });
      return response.data.valid;
    } catch (error) {
      console.error('Token verification failed:', error);
      return false;
    }
  }
};