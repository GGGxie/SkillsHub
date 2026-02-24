const API_BASE = '/api'

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const token = localStorage.getItem('token')
  const headers: Record<string, string> = {
    'Content-Type': 'application/json',
    ...(options?.headers as Record<string, string>),
  }
  if (token) {
    headers['Authorization'] = `Bearer ${token}`
  }

  const res = await fetch(`${API_BASE}${url}`, { ...options, headers })
  if (!res.ok) {
    const err = await res.json().catch(() => ({ error: 'Request failed' }))
    throw new Error(err.error || 'Request failed')
  }
  return res.json()
}

export interface Skill {
  id: number
  title: string
  title_zh: string
  description: string
  desc_zh: string
  category: string
  icon: string
  image: string
  author_id: number
  author_name: string
  author_avatar: string
  content: string
  content_zh: string
  tags: string
  skill_type: string
  featured: boolean
  likes: number
  views: number
  created_at: string
  updated_at: string
}

export interface Article {
  id: number
  title: string
  title_zh: string
  description: string
  desc_zh: string
  category: string
  content: string
  content_zh: string
  author_id: number
  author_name: string
  views: number
  created_at: string
}

export interface Comment {
  id: number
  skill_id: number
  user_id: number
  user_name: string
  avatar: string
  content: string
  created_at: string
}

export interface User {
  id: number
  google_id: string
  email: string
  name: string
  avatar: string
  created_at: string
}

export interface PaginatedResponse<T> {
  data: T[]
  total: number
  page: number
  page_size: number
  total_pages: number
}

export const api = {
  skills: {
    list: (params: { page?: number; sort?: string; category?: string; search?: string }) => {
      const q = new URLSearchParams()
      if (params.page) q.set('page', String(params.page))
      if (params.sort) q.set('sort', params.sort)
      if (params.category) q.set('category', params.category)
      if (params.search) q.set('search', params.search)
      return request<PaginatedResponse<Skill>>(`/skills?${q}`)
    },
    featured: () => request<Skill[]>('/skills/featured'),
    get: (id: number | string) => request<Skill>(`/skills/${id}`),
    create: (data: Partial<Skill>) => request<{ id: number }>('/skills', { method: 'POST', body: JSON.stringify(data) }),
    like: (id: number | string) => request<{ liked: boolean }>(`/skills/${id}/like`, { method: 'POST' }),
    categories: () => request<string[]>('/skills/categories'),
    comments: (id: number | string) => request<Comment[]>(`/skills/${id}/comments`),
    addComment: (id: number | string, content: string) =>
      request<{ id: number }>(`/skills/${id}/comments`, { method: 'POST', body: JSON.stringify({ content }) }),
  },
  articles: {
    list: (category?: string) => {
      const q = category ? `?category=${category}` : ''
      return request<Article[]>(`/articles${q}`)
    },
    get: (id: number | string) => request<Article>(`/articles/${id}`),
  },
  auth: {
    googleToken: (credential: string) =>
      request<{ token: string; user: User }>('/auth/google/token', { method: 'POST', body: JSON.stringify({ credential }) }),
    me: () => request<User>('/auth/me'),
  },
}
