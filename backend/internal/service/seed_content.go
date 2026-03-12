package service

var cb = "```"

type seedSkill struct {
	title, titleZH, desc, descZH, category, icon, tags, skillType string
	content, contentZH                                             string
	featured                                                       int
	likes, views                                                   int
}

var seedSkills = []seedSkill{
	{
		title: "Code Review Assistant", titleZH: "代码审查助手",
		desc: "Automated code review with best practices and security checks", descZH: "自动化代码审查，包含最佳实践和安全检查",
		category: "Coding & Development", icon: "🔍", tags: "code,review,security,best-practices", skillType: "ai-powered", featured: 1, likes: 7, views: 48,
		content: "# Code Review Assistant\n\n" +
			"Perform systematic code reviews covering security, correctness, performance, and maintainability.\n\n" +
			"## Review Checklist\n\n" +
			"### Security\n" +
			"- [ ] No SQL injection (use parameterized queries)\n" +
			"- [ ] No XSS vulnerabilities (sanitize user input)\n" +
			"- [ ] No hardcoded secrets, API keys, or credentials\n" +
			"- [ ] Input validation on all user-facing endpoints\n" +
			"- [ ] Authentication & authorization checks in place\n\n" +
			"### Correctness\n" +
			"- [ ] Edge cases handled: null/undefined, empty arrays, boundary values\n" +
			"- [ ] Error handling is comprehensive (no swallowed errors)\n" +
			"- [ ] Async/await and promise chains are correct\n" +
			"- [ ] Race conditions considered in concurrent code\n\n" +
			"### Performance\n" +
			"- [ ] No N+1 query patterns\n" +
			"- [ ] Unnecessary re-renders avoided (React: useMemo, useCallback)\n" +
			"- [ ] No memory leaks (event listeners, subscriptions cleaned up)\n" +
			"- [ ] Database queries use proper indexes\n\n" +
			"### Maintainability\n" +
			"- [ ] Functions under 30 lines, single responsibility\n" +
			"- [ ] No code duplication (DRY principle)\n" +
			"- [ ] Naming is clear and consistent\n" +
			"- [ ] Complex logic has explanatory comments\n\n" +
			"## Output Format\n\n" +
			"Rate each issue found:\n" +
			"- 🔴 **Critical**: Must fix — security vulnerabilities, data loss, crashes\n" +
			"- 🟡 **Warning**: Should fix — performance issues, potential bugs\n" +
			"- 🟢 **Suggestion**: Nice to have — readability, style improvements\n\n" +
			"## Example\n\n" +
			cb + "diff\n" +
			"- password = request.form['password']\n" +
			"- db.execute(f\"SELECT * FROM users WHERE pass='{password}'\")\n" +
			"+ password_hash = hash_password(request.form['password'])\n" +
			"+ db.execute(\"SELECT * FROM users WHERE pass=?\", (password_hash,))\n" +
			cb + "\n\n" +
			"🔴 **Critical**: SQL injection + plaintext password storage. Use parameterized queries and hash passwords.\n",
		contentZH: "# 代码审查助手\n\n" +
			"系统性地进行代码审查，涵盖安全性、正确性、性能和可维护性。\n\n" +
			"## 审查清单\n\n" +
			"### 安全性\n" +
			"- [ ] 无 SQL 注入（使用参数化查询）\n" +
			"- [ ] 无 XSS 漏洞（对用户输入进行转义）\n" +
			"- [ ] 无硬编码的密钥、API Key 或凭证\n" +
			"- [ ] 所有用户接口都有输入验证\n" +
			"- [ ] 认证和授权检查完备\n\n" +
			"### 正确性\n" +
			"- [ ] 边界情况已处理：null/undefined、空数组、边界值\n" +
			"- [ ] 错误处理完善（不吞掉异常）\n" +
			"- [ ] async/await 和 Promise 链正确\n" +
			"- [ ] 并发代码考虑了竞态条件\n\n" +
			"### 性能\n" +
			"- [ ] 无 N+1 查询\n" +
			"- [ ] 避免不必要的重新渲染（React: useMemo, useCallback）\n" +
			"- [ ] 无内存泄漏（事件监听器、订阅已清理）\n" +
			"- [ ] 数据库查询使用了合适的索引\n\n" +
			"### 可维护性\n" +
			"- [ ] 函数不超过 30 行，职责单一\n" +
			"- [ ] 无代码重复（DRY 原则）\n" +
			"- [ ] 命名清晰一致\n" +
			"- [ ] 复杂逻辑有解释性注释\n\n" +
			"## 输出格式\n\n" +
			"每个问题的严重程度：\n" +
			"- 🔴 **严重**：必须修复 — 安全漏洞、数据丢失、崩溃\n" +
			"- 🟡 **警告**：建议修复 — 性能问题、潜在 bug\n" +
			"- 🟢 **建议**：可选改进 — 可读性、风格优化\n\n" +
			"## 示例\n\n" +
			cb + "diff\n" +
			"- password = request.form['password']\n" +
			"- db.execute(f\"SELECT * FROM users WHERE pass='{password}'\")\n" +
			"+ password_hash = hash_password(request.form['password'])\n" +
			"+ db.execute(\"SELECT * FROM users WHERE pass=?\", (password_hash,))\n" +
			cb + "\n\n" +
			"🔴 **严重**：SQL 注入 + 明文密码存储。使用参数化查询并对密码加密。\n",
	},
	{
		title: "Blog Post Writer", titleZH: "博客文章写手",
		desc: "Generate well-structured blog posts with SEO optimization", descZH: "生成结构良好的博客文章，包含SEO优化",
		category: "Writing & Content", icon: "✍️", tags: "writing,blog,seo,content-marketing", skillType: "general", featured: 1, likes: 6, views: 38,
		content: "# Blog Post Writer\n\n" +
			"Generate well-structured, SEO-optimized blog posts that engage readers and rank well.\n\n" +
			"## Article Structure\n\n" +
			"### Title (< 60 characters)\n" +
			"- Include primary keyword naturally\n" +
			"- Use numbers, \"How to\", or power words for higher CTR\n" +
			"- Examples: \"7 Proven Ways to...\", \"The Complete Guide to...\"\n\n" +
			"### Introduction (100-150 words)\n" +
			"1. Open with a hook: a surprising stat, question, or pain point\n" +
			"2. State the problem the reader faces\n" +
			"3. Preview what they'll learn (the promise)\n" +
			"4. Include the primary keyword in the first paragraph\n\n" +
			"### Body (3-5 sections with H2/H3 headings)\n" +
			"- Each section addresses one key point\n" +
			"- Paragraphs: 3-4 sentences max\n" +
			"- Use bullet points and numbered lists for scannability\n" +
			"- Include examples, data, or case studies\n" +
			"- Add transition sentences between sections\n\n" +
			"### Conclusion\n" +
			"- Summarize 3 key takeaways\n" +
			"- Include a clear call to action\n" +
			"- End with an engaging question or forward-looking statement\n\n" +
			"## SEO Checklist\n\n" +
			"- [ ] Primary keyword in: title, H1, first paragraph, at least 2 H2s\n" +
			"- [ ] Meta description: < 160 characters, includes keyword, has CTA\n" +
			"- [ ] 2-3 internal links + 1-2 external authority links\n" +
			"- [ ] Images with descriptive alt text\n" +
			"- [ ] URL slug is short and keyword-rich\n" +
			"- [ ] Reading level: Grade 6-8 (use short sentences)\n\n" +
			"## Tone Guidelines\n\n" +
			"| Audience | Tone | Example |\n" +
			"|----------|------|---------|\n" +
			"| Developers | Technical, concise | \"Implement caching with Redis to reduce latency by 40%\" |\n" +
			"| Business | Professional, benefit-focused | \"This approach saves teams 10+ hours per week\" |\n" +
			"| General | Friendly, conversational | \"Think of it like organizing your digital toolbox\" |\n",
		contentZH: "# 博客文章写手\n\n" +
			"生成结构合理、SEO 优化良好的博客文章，吸引读者并获得好的搜索排名。\n\n" +
			"## 文章结构\n\n" +
			"### 标题（< 60 字符）\n" +
			"- 自然融入主关键词\n" +
			"- 使用数字、\"如何\"、或有力的词语提升点击率\n" +
			"- 示例：\"7个经验证的方法...\"、\"完全指南...\" \n\n" +
			"### 引言（100-150 字）\n" +
			"1. 用令人惊讶的数据、问题或痛点开场\n" +
			"2. 阐述读者面临的问题\n" +
			"3. 预告他们将学到什么\n" +
			"4. 在首段自然包含主关键词\n\n" +
			"### 正文（3-5 个章节，使用 H2/H3 标题）\n" +
			"- 每个章节聚焦一个要点\n" +
			"- 段落：最多 3-4 句\n" +
			"- 使用项目符号和编号列表提高可扫描性\n" +
			"- 包含示例、数据或案例研究\n\n" +
			"### 结论\n" +
			"- 总结 3 个核心要点\n" +
			"- 包含明确的行动号召\n" +
			"- 以引人深思的问题或展望结尾\n\n" +
			"## SEO 清单\n\n" +
			"- [ ] 主关键词出现在：标题、H1、首段、至少 2 个 H2\n" +
			"- [ ] Meta 描述：< 160 字符，包含关键词，有 CTA\n" +
			"- [ ] 2-3 个内链 + 1-2 个权威外链\n" +
			"- [ ] 图片有描述性 alt 文本\n" +
			"- [ ] URL 简短且包含关键词\n\n" +
			"## 语气指南\n\n" +
			"| 受众 | 语气 | 示例 |\n" +
			"|------|------|------|\n" +
			"| 开发者 | 技术性、简洁 | \"使用 Redis 缓存可将延迟降低 40%\" |\n" +
			"| 商务 | 专业、关注效益 | \"此方案每周为团队节省 10+ 小时\" |\n" +
			"| 大众 | 友好、对话式 | \"把它想象成整理你的数字工具箱\" |\n",
	},
	{
		title: "Data Analysis Pipeline", titleZH: "数据分析流水线",
		desc: "Automated data analysis with visualization and reporting", descZH: "自动化数据分析，包含可视化和报告",
		category: "Data & Analytics", icon: "📊", tags: "data,analysis,visualization,python,pandas", skillType: "ai-powered", featured: 1, likes: 5, views: 33,
		content: "# Data Analysis Pipeline\n\n" +
			"A step-by-step workflow for loading, cleaning, analyzing, and visualizing data with Python.\n\n" +
			"## Workflow\n\n" +
			"### Step 1: Load & Inspect\n\n" +
			cb + "python\n" +
			"import pandas as pd\n\n" +
			"df = pd.read_csv('data.csv')\n" +
			"print(df.shape, df.dtypes)\n" +
			"print(df.describe())\n" +
			"print(df.isnull().sum())\n" +
			cb + "\n\n" +
			"### Step 2: Clean\n" +
			"- Drop duplicates: `df.drop_duplicates(inplace=True)`\n" +
			"- Handle missing values: fill with median/mode or drop rows\n" +
			"- Fix data types: `df['date'] = pd.to_datetime(df['date'])`\n" +
			"- Remove outliers using IQR or z-score methods\n\n" +
			"### Step 3: Analyze\n\n" +
			cb + "python\n" +
			"# Grouping and aggregation\n" +
			"summary = df.groupby('category').agg(\n" +
			"    count=('id', 'count'),\n" +
			"    avg_value=('value', 'mean'),\n" +
			"    total=('value', 'sum')\n" +
			").sort_values('total', ascending=False)\n\n" +
			"# Correlation analysis\n" +
			"corr_matrix = df.select_dtypes(include='number').corr()\n" +
			cb + "\n\n" +
			"### Step 4: Visualize\n\n" +
			cb + "python\n" +
			"import matplotlib.pyplot as plt\n" +
			"import seaborn as sns\n\n" +
			"fig, axes = plt.subplots(1, 3, figsize=(15, 5))\n" +
			"sns.barplot(data=summary.reset_index(), x='category', y='total', ax=axes[0])\n" +
			"sns.heatmap(corr_matrix, annot=True, cmap='coolwarm', ax=axes[1])\n" +
			"df['value'].hist(bins=30, ax=axes[2])\n" +
			"plt.tight_layout()\n" +
			"plt.savefig('analysis.png', dpi=150)\n" +
			cb + "\n\n" +
			"### Step 5: Report\n" +
			"Generate a summary report with:\n" +
			"- Key findings (top 3-5 insights)\n" +
			"- Statistical significance of results\n" +
			"- Visualizations with clear labels\n" +
			"- Actionable recommendations\n",
		contentZH: "# 数据分析流水线\n\n" +
			"使用 Python 进行数据加载、清洗、分析和可视化的完整工作流。\n\n" +
			"## 工作流程\n\n" +
			"### 步骤 1：加载和检查\n\n" +
			cb + "python\n" +
			"import pandas as pd\n\n" +
			"df = pd.read_csv('data.csv')\n" +
			"print(df.shape, df.dtypes)\n" +
			"print(df.describe())\n" +
			"print(df.isnull().sum())\n" +
			cb + "\n\n" +
			"### 步骤 2：清洗\n" +
			"- 删除重复项：`df.drop_duplicates(inplace=True)`\n" +
			"- 处理缺失值：用中位数/众数填充或删除行\n" +
			"- 修正数据类型：`df['date'] = pd.to_datetime(df['date'])`\n" +
			"- 使用 IQR 或 z-score 方法移除异常值\n\n" +
			"### 步骤 3：分析\n\n" +
			cb + "python\n" +
			"summary = df.groupby('category').agg(\n" +
			"    count=('id', 'count'),\n" +
			"    avg_value=('value', 'mean'),\n" +
			"    total=('value', 'sum')\n" +
			").sort_values('total', ascending=False)\n" +
			cb + "\n\n" +
			"### 步骤 4：可视化\n\n" +
			cb + "python\n" +
			"import matplotlib.pyplot as plt\n" +
			"import seaborn as sns\n\n" +
			"fig, axes = plt.subplots(1, 3, figsize=(15, 5))\n" +
			"sns.barplot(data=summary.reset_index(), x='category', y='total', ax=axes[0])\n" +
			"sns.heatmap(corr_matrix, annot=True, cmap='coolwarm', ax=axes[1])\n" +
			"df['value'].hist(bins=30, ax=axes[2])\n" +
			"plt.tight_layout()\n" +
			"plt.savefig('analysis.png', dpi=150)\n" +
			cb + "\n\n" +
			"### 步骤 5：生成报告\n" +
			"生成包含以下内容的总结报告：\n" +
			"- 关键发现（前 3-5 个洞察）\n" +
			"- 结果的统计显著性\n" +
			"- 带有清晰标签的可视化图表\n" +
			"- 可操作的建议\n",
	},
	{
		title: "API Integration Helper", titleZH: "API 集成助手",
		desc: "Simplify REST API integration with auto-generated client code", descZH: "简化 REST API 集成，自动生成客户端代码",
		category: "Coding & Development", icon: "🔗", tags: "api,integration,rest,http,typescript", skillType: "general", featured: 0, likes: 4, views: 29,
		content: "# API Integration Helper\n\n" +
			"Generate type-safe API client code from endpoint specifications.\n\n" +
			"## Input Format\n\n" +
			"Describe the API endpoint:\n" +
			"- Method and URL\n" +
			"- Request body / query parameters\n" +
			"- Expected response shape\n" +
			"- Authentication method\n\n" +
			"## Generated Output\n\n" +
			"### TypeScript API Client\n\n" +
			cb + "typescript\n" +
			"interface ApiConfig {\n" +
			"  baseURL: string;\n" +
			"  token?: string;\n" +
			"  timeout?: number;\n" +
			"}\n\n" +
			"async function apiRequest<T>(\n" +
			"  config: ApiConfig,\n" +
			"  method: string,\n" +
			"  path: string,\n" +
			"  body?: unknown\n" +
			"): Promise<T> {\n" +
			"  const res = await fetch(`${config.baseURL}${path}`, {\n" +
			"    method,\n" +
			"    headers: {\n" +
			"      'Content-Type': 'application/json',\n" +
			"      ...(config.token && { Authorization: `Bearer ${config.token}` }),\n" +
			"    },\n" +
			"    body: body ? JSON.stringify(body) : undefined,\n" +
			"    signal: AbortSignal.timeout(config.timeout ?? 10000),\n" +
			"  });\n" +
			"  if (!res.ok) {\n" +
			"    const error = await res.json().catch(() => ({}));\n" +
			"    throw new ApiError(res.status, error.message ?? res.statusText);\n" +
			"  }\n" +
			"  return res.json();\n" +
			"}\n" +
			cb + "\n\n" +
			"## Best Practices\n\n" +
			"- Always define TypeScript interfaces for request/response\n" +
			"- Implement retry logic with exponential backoff for 5xx errors\n" +
			"- Use AbortController for request cancellation\n" +
			"- Log request/response for debugging (redact sensitive data)\n" +
			"- Handle rate limiting (429) with automatic retry after delay\n",
		contentZH: "# API 集成助手\n\n" +
			"根据接口规格生成类型安全的 API 客户端代码。\n\n" +
			"## 输入格式\n\n" +
			"描述 API 端点：\n" +
			"- 请求方法和 URL\n" +
			"- 请求体/查询参数\n" +
			"- 预期响应结构\n" +
			"- 认证方式\n\n" +
			"## 生成输出\n\n" +
			"### TypeScript API 客户端\n\n" +
			cb + "typescript\n" +
			"interface ApiConfig {\n" +
			"  baseURL: string;\n" +
			"  token?: string;\n" +
			"  timeout?: number;\n" +
			"}\n\n" +
			"async function apiRequest<T>(\n" +
			"  config: ApiConfig,\n" +
			"  method: string,\n" +
			"  path: string,\n" +
			"  body?: unknown\n" +
			"): Promise<T> {\n" +
			"  const res = await fetch(`${config.baseURL}${path}`, {\n" +
			"    method,\n" +
			"    headers: {\n" +
			"      'Content-Type': 'application/json',\n" +
			"      ...(config.token && { Authorization: `Bearer ${config.token}` }),\n" +
			"    },\n" +
			"    body: body ? JSON.stringify(body) : undefined,\n" +
			"  });\n" +
			"  if (!res.ok) throw new Error(res.statusText);\n" +
			"  return res.json();\n" +
			"}\n" +
			cb + "\n\n" +
			"## 最佳实践\n\n" +
			"- 为请求和响应定义 TypeScript 接口\n" +
			"- 对 5xx 错误实现指数退避重试\n" +
			"- 使用 AbortController 实现请求取消\n" +
			"- 记录请求/响应日志（脱敏处理）\n" +
			"- 处理速率限制（429），自动延迟重试\n",
	},
	{
		title: "Email Template Designer", titleZH: "邮件模板设计师",
		desc: "Create responsive email templates with dynamic content", descZH: "创建响应式邮件模板，支持动态内容",
		category: "Design & Creative", icon: "📧", tags: "email,template,design,html,responsive", skillType: "general", featured: 0, likes: 4, views: 26,
		content: "# Email Template Designer\n\n" +
			"Create responsive, cross-client compatible HTML email templates.\n\n" +
			"## Design Rules\n\n" +
			"Email HTML is NOT like web HTML. Follow these constraints:\n\n" +
			"| Do | Don't |\n" +
			"|-----|-------|\n" +
			"| Use `<table>` for layout | Use `<div>` for layout |\n" +
			"| Inline CSS styles | Use `<link>` or `<style>` |\n" +
			"| Use `width` attribute on tables | Use CSS `max-width` alone |\n" +
			"| Use web-safe fonts + fallbacks | Use custom web fonts |\n" +
			"| 600px max width | Fluid full-width layouts |\n\n" +
			"## Template Structure\n\n" +
			cb + "html\n" +
			"<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" bgcolor=\"#f4f4f4\">\n" +
			"  <tr><td align=\"center\">\n" +
			"    <table width=\"600\" cellpadding=\"0\" cellspacing=\"0\" bgcolor=\"#ffffff\">\n" +
			"      <!-- Header with logo -->\n" +
			"      <tr><td style=\"padding:20px;text-align:center\">\n" +
			"        <img src=\"logo.png\" width=\"150\" alt=\"Logo\">\n" +
			"      </td></tr>\n" +
			"      <!-- Content -->\n" +
			"      <tr><td style=\"padding:20px 30px;font-family:Arial,sans-serif\">\n" +
			"        <h1 style=\"color:#333;font-size:24px\">{{title}}</h1>\n" +
			"        <p style=\"color:#666;font-size:16px;line-height:1.5\">{{body}}</p>\n" +
			"      </td></tr>\n" +
			"      <!-- CTA Button -->\n" +
			"      <tr><td align=\"center\" style=\"padding:20px\">\n" +
			"        <a href=\"{{cta_url}}\" style=\"background:#007bff;color:#fff;\n" +
			"          padding:12px 30px;text-decoration:none;border-radius:4px;\n" +
			"          display:inline-block\">{{cta_text}}</a>\n" +
			"      </td></tr>\n" +
			"    </table>\n" +
			"  </td></tr>\n" +
			"</table>\n" +
			cb + "\n\n" +
			"## Testing Checklist\n\n" +
			"- [ ] Renders correctly in Gmail, Outlook, Apple Mail\n" +
			"- [ ] Images have alt text and fallback background colors\n" +
			"- [ ] Links are tracked and working\n" +
			"- [ ] Plain text version is provided\n" +
			"- [ ] Responsive on mobile (< 480px)\n",
		contentZH: "# 邮件模板设计师\n\n" +
			"创建跨客户端兼容的响应式 HTML 邮件模板。\n\n" +
			"## 设计规则\n\n" +
			"邮件 HTML 与网页 HTML 不同，需遵循以下约束：\n\n" +
			"| 应该 | 不应该 |\n" +
			"|------|--------|\n" +
			"| 使用 `<table>` 布局 | 使用 `<div>` 布局 |\n" +
			"| 内联 CSS 样式 | 使用外部样式表 |\n" +
			"| 最大宽度 600px | 全宽流式布局 |\n" +
			"| 使用网页安全字体 | 使用自定义字体 |\n\n" +
			"## 模板结构\n\n" +
			cb + "html\n" +
			"<table width=\"100%\" cellpadding=\"0\" cellspacing=\"0\" bgcolor=\"#f4f4f4\">\n" +
			"  <tr><td align=\"center\">\n" +
			"    <table width=\"600\" cellpadding=\"0\" cellspacing=\"0\" bgcolor=\"#ffffff\">\n" +
			"      <!-- 头部 Logo -->\n" +
			"      <tr><td style=\"padding:20px;text-align:center\">\n" +
			"        <img src=\"logo.png\" width=\"150\" alt=\"Logo\">\n" +
			"      </td></tr>\n" +
			"      <!-- 内容 -->\n" +
			"      <tr><td style=\"padding:20px 30px;font-family:Arial,sans-serif\">\n" +
			"        <h1 style=\"color:#333;font-size:24px\">{{title}}</h1>\n" +
			"        <p style=\"color:#666;font-size:16px;line-height:1.5\">{{body}}</p>\n" +
			"      </td></tr>\n" +
			"      <!-- CTA 按钮 -->\n" +
			"      <tr><td align=\"center\" style=\"padding:20px\">\n" +
			"        <a href=\"{{cta_url}}\" style=\"background:#007bff;color:#fff;\n" +
			"          padding:12px 30px;text-decoration:none;border-radius:4px;\n" +
			"          display:inline-block\">{{cta_text}}</a>\n" +
			"      </td></tr>\n" +
			"    </table>\n" +
			"  </td></tr>\n" +
			"</table>\n" +
			cb + "\n\n" +
			"## 测试清单\n\n" +
			"- [ ] 在 Gmail、Outlook、Apple Mail 中正确渲染\n" +
			"- [ ] 图片有 alt 文本和回退背景色\n" +
			"- [ ] 链接可追踪且有效\n" +
			"- [ ] 提供纯文本版本\n" +
			"- [ ] 移动端响应式（< 480px）\n",
	},
	{
		title: "Test Case Generator", titleZH: "测试用例生成器",
		desc: "Auto-generate comprehensive test cases from specifications", descZH: "根据规格说明自动生成全面的测试用例",
		category: "Coding & Development", icon: "🧪", tags: "testing,automation,quality,jest,pytest", skillType: "ai-powered", featured: 1, likes: 7, views: 42,
		content: "# Test Case Generator\n\n" +
			"Generate comprehensive test cases from function signatures, API specs, or user stories.\n\n" +
			"## Test Categories\n\n" +
			"For any given function or endpoint, generate tests covering:\n\n" +
			"1. **Happy Path** — Normal expected behavior\n" +
			"2. **Edge Cases** — Boundary values, empty inputs, max values\n" +
			"3. **Error Cases** — Invalid inputs, missing required fields\n" +
			"4. **Security** — Injection, unauthorized access, data leaks\n" +
			"5. **Performance** — Large datasets, concurrent requests\n\n" +
			"## Template (Jest/TypeScript)\n\n" +
			cb + "typescript\n" +
			"describe('UserService.createUser', () => {\n" +
			"  // Happy path\n" +
			"  it('should create user with valid data', async () => {\n" +
			"    const user = await service.createUser({ name: 'Alice', email: 'a@b.com' });\n" +
			"    expect(user.id).toBeDefined();\n" +
			"    expect(user.name).toBe('Alice');\n" +
			"  });\n\n" +
			"  // Edge cases\n" +
			"  it('should handle name at max length (255 chars)', async () => { ... });\n" +
			"  it('should trim whitespace from email', async () => { ... });\n\n" +
			"  // Error cases\n" +
			"  it('should reject duplicate email', async () => {\n" +
			"    await service.createUser({ name: 'A', email: 'dup@b.com' });\n" +
			"    await expect(service.createUser({ name: 'B', email: 'dup@b.com' }))\n" +
			"      .rejects.toThrow('Email already exists');\n" +
			"  });\n" +
			"  it('should reject invalid email format', async () => { ... });\n" +
			"  it('should reject empty name', async () => { ... });\n" +
			"});\n" +
			cb + "\n\n" +
			"## Template (Pytest)\n\n" +
			cb + "python\n" +
			"import pytest\n\n" +
			"class TestCreateUser:\n" +
			"    def test_creates_user_with_valid_data(self, db):\n" +
			"        user = create_user(name='Alice', email='a@b.com')\n" +
			"        assert user.id is not None\n\n" +
			"    def test_rejects_duplicate_email(self, db):\n" +
			"        create_user(name='A', email='dup@b.com')\n" +
			"        with pytest.raises(ValueError, match='already exists'):\n" +
			"            create_user(name='B', email='dup@b.com')\n\n" +
			"    @pytest.mark.parametrize('email', ['', 'invalid', '@no-user.com', 'no@'])\n" +
			"    def test_rejects_invalid_email(self, db, email):\n" +
			"        with pytest.raises(ValueError):\n" +
			"            create_user(name='Test', email=email)\n" +
			cb + "\n\n" +
			"## Coverage Strategy\n\n" +
			"Aim for: 80%+ line coverage, 100% branch coverage on critical paths (auth, payments, data mutations).\n",
		contentZH: "# 测试用例生成器\n\n" +
			"根据函数签名、API 规格或用户故事生成全面的测试用例。\n\n" +
			"## 测试类别\n\n" +
			"对任何函数或接口，生成以下测试：\n\n" +
			"1. **正常路径** — 预期的正常行为\n" +
			"2. **边界情况** — 边界值、空输入、最大值\n" +
			"3. **错误情况** — 无效输入、缺少必填字段\n" +
			"4. **安全性** — 注入攻击、未授权访问\n" +
			"5. **性能** — 大数据集、并发请求\n\n" +
			"## 模板（Jest/TypeScript）\n\n" +
			cb + "typescript\n" +
			"describe('UserService.createUser', () => {\n" +
			"  // 正常路径\n" +
			"  it('should create user with valid data', async () => {\n" +
			"    const user = await service.createUser({ name: 'Alice', email: 'a@b.com' });\n" +
			"    expect(user.id).toBeDefined();\n" +
			"    expect(user.name).toBe('Alice');\n" +
			"  });\n\n" +
			"  // 边界情况\n" +
			"  it('should handle name at max length (255 chars)', async () => { ... });\n" +
			"  it('should trim whitespace from email', async () => { ... });\n\n" +
			"  // 错误情况\n" +
			"  it('should reject duplicate email', async () => {\n" +
			"    await service.createUser({ name: 'A', email: 'dup@b.com' });\n" +
			"    await expect(service.createUser({ name: 'B', email: 'dup@b.com' }))\n" +
			"      .rejects.toThrow('Email already exists');\n" +
			"  });\n" +
			"  it('should reject invalid email format', async () => { ... });\n" +
			"  it('should reject empty name', async () => { ... });\n" +
			"});\n" +
			cb + "\n\n" +
			"## 模板（Pytest）\n\n" +
			cb + "python\n" +
			"import pytest\n\n" +
			"class TestCreateUser:\n" +
			"    def test_creates_user_with_valid_data(self, db):\n" +
			"        user = create_user(name='Alice', email='a@b.com')\n" +
			"        assert user.id is not None\n\n" +
			"    def test_rejects_duplicate_email(self, db):\n" +
			"        create_user(name='A', email='dup@b.com')\n" +
			"        with pytest.raises(ValueError, match='already exists'):\n" +
			"            create_user(name='B', email='dup@b.com')\n\n" +
			"    @pytest.mark.parametrize('email', ['', 'invalid', '@no-user.com', 'no@'])\n" +
			"    def test_rejects_invalid_email(self, db, email):\n" +
			"        with pytest.raises(ValueError):\n" +
			"            create_user(name='Test', email=email)\n" +
			cb + "\n\n" +
			"## 覆盖策略\n\n" +
			"目标：80%+ 行覆盖率，关键路径（认证、支付、数据变更）100% 分支覆盖。\n",
	},
	{
		title: "Translation Assistant", titleZH: "翻译助手",
		desc: "Multi-language translation with context awareness", descZH: "多语言翻译，支持上下文感知",
		category: "Writing & Content", icon: "🌐", tags: "translation,language,i18n,localization", skillType: "ai-powered", featured: 0, likes: 3, views: 20,
		content: "# Translation Assistant\n\n" +
			"Context-aware translation that preserves meaning, tone, and cultural nuance.\n\n" +
			"## Translation Principles\n\n" +
			"1. **Meaning over words** — Translate the intent, not word-for-word\n" +
			"2. **Preserve tone** — Formal stays formal, casual stays casual\n" +
			"3. **Cultural adaptation** — Adjust idioms, metaphors, and references\n" +
			"4. **Technical accuracy** — Keep technical terms consistent\n" +
			"5. **Natural flow** — Output should read as native text\n\n" +
			"## Context-Aware Rules\n\n" +
			"| Context | Rule | Example |\n" +
			"|---------|------|---------|\n" +
			"| UI strings | Keep short, use standard UX terms | \"Submit\" → \"提交\" (not \"递交\") |\n" +
			"| Documentation | Formal, precise | Use \"您\" not \"你\" in Chinese |\n" +
			"| Marketing | Persuasive, locally adapted | Adapt CTAs to local conventions |\n" +
			"| Legal | Exact meaning, no interpretation | Preserve legal terminology |\n" +
			"| Chat/casual | Natural, colloquial | Use conversational language |\n\n" +
			"## For Developers (i18n)\n\n" +
			"When translating JSON locale files:\n\n" +
			cb + "json\n" +
			"{\n" +
			"  \"greeting\": \"Hello, {{name}}!\",\n" +
			"  \"items_count\": \"{{count}} item | {{count}} items\",\n" +
			"  \"error.network\": \"Connection failed. Please try again.\"\n" +
			"}\n" +
			cb + "\n\n" +
			"Rules:\n" +
			"- Preserve all `{{variables}}` and placeholders exactly\n" +
			"- Handle pluralization rules for target language\n" +
			"- Keep JSON keys unchanged\n" +
			"- Maintain consistent terminology across all keys\n",
		contentZH: "# 翻译助手\n\n" +
			"上下文感知的翻译，保留含义、语气和文化细微差别。\n\n" +
			"## 翻译原则\n\n" +
			"1. **达意优先** — 翻译意图，而非逐字翻译\n" +
			"2. **保持语气** — 正式的保持正式，轻松的保持轻松\n" +
			"3. **文化适配** — 调整习语、比喻和文化引用\n" +
			"4. **术语一致** — 技术术语保持统一\n" +
			"5. **自然流畅** — 译文应读起来像母语文本\n\n" +
			"## 上下文规则\n\n" +
			"| 场景 | 规则 | 示例 |\n" +
			"|------|------|------|\n" +
			"| UI 字符串 | 简洁，使用标准 UX 用语 | \"Submit\" → \"提交\"（而非\"递交\"） |\n" +
			"| 文档 | 正式、精确 | 中文使用\"您\"而非\"你\" |\n" +
			"| 营销 | 有说服力，本地化适配 | 调整 CTA 适配本地习惯 |\n" +
			"| 法律 | 精确含义，不做引申 | 保留法律术语 |\n" +
			"| 聊天/日常 | 自然、口语化 | 使用日常对话语言 |\n\n" +
			"## 开发者指南（i18n）\n\n" +
			"翻译 JSON 语言包时：\n\n" +
			cb + "json\n" +
			"{\n" +
			"  \"greeting\": \"Hello, {{name}}!\",\n" +
			"  \"items_count\": \"{{count}} item | {{count}} items\",\n" +
			"  \"error.network\": \"Connection failed. Please try again.\"\n" +
			"}\n" +
			cb + "\n\n" +
			"规则：\n" +
			"- 保留所有 `{{变量}}` 和占位符\n" +
			"- 处理目标语言的复数规则\n" +
			"- JSON key 保持不变\n" +
			"- 所有 key 之间术语一致\n",
	},
	{
		title: "Task Automation Bot", titleZH: "任务自动化机器人",
		desc: "Automate repetitive tasks with customizable workflows", descZH: "自动化重复任务，支持可自定义工作流",
		category: "Automation", icon: "🤖", tags: "automation,workflow,scripting,productivity", skillType: "general", featured: 0, likes: 6, views: 36,
		content: "# Task Automation Bot\n\n" +
			"Design and implement automation workflows for repetitive tasks.\n\n" +
			"## Workflow Design Pattern\n\n" +
			"### 1. Identify the Task\n" +
			"- What triggers it? (time-based, event-based, manual)\n" +
			"- What are the inputs and outputs?\n" +
			"- What are the failure modes?\n\n" +
			"### 2. Design the Pipeline\n\n" +
			cb + "yaml\n" +
			"# workflow.yml\n" +
			"name: daily-report\n" +
			"trigger:\n" +
			"  schedule: '0 9 * * *'  # Every day at 9 AM\n" +
			"steps:\n" +
			"  - name: fetch-data\n" +
			"    action: http.get\n" +
			"    url: https://api.example.com/metrics\n" +
			"  - name: process\n" +
			"    action: transform\n" +
			"    input: ${{ steps.fetch-data.output }}\n" +
			"  - name: notify\n" +
			"    action: slack.send\n" +
			"    channel: '#reports'\n" +
			"    message: ${{ steps.process.output }}\n" +
			"    on_failure: email.alert\n" +
			cb + "\n\n" +
			"### 3. Error Handling\n\n" +
			"Every automation must handle:\n" +
			"- **Retry logic**: 3 attempts with exponential backoff\n" +
			"- **Timeout**: Set max execution time\n" +
			"- **Alerting**: Notify on failure (Slack, email)\n" +
			"- **Idempotency**: Safe to re-run without side effects\n\n" +
			"### 4. Common Automations\n\n" +
			"| Task | Tool | Trigger |\n" +
			"|------|------|---------|\n" +
			"| File backup | rsync + cron | Schedule |\n" +
			"| DB cleanup | SQL script + cron | Schedule |\n" +
			"| Deploy | GitHub Actions | Git push |\n" +
			"| Log rotation | logrotate | Daily |\n" +
			"| Health check | curl + cron | Every 5 min |\n",
		contentZH: "# 任务自动化机器人\n\n" +
			"设计和实现重复任务的自动化工作流。\n\n" +
			"## 工作流设计\n\n" +
			"### 1. 识别任务\n" +
			"- 什么触发它？（定时、事件驱动、手动）\n" +
			"- 输入和输出是什么？\n" +
			"- 可能的失败模式？\n\n" +
			"### 2. 设计流水线\n\n" +
			cb + "yaml\n" +
			"# workflow.yml\n" +
			"name: daily-report\n" +
			"trigger:\n" +
			"  schedule: '0 9 * * *'  # 每天早上 9 点\n" +
			"steps:\n" +
			"  - name: fetch-data\n" +
			"    action: http.get\n" +
			"    url: https://api.example.com/metrics\n" +
			"  - name: process\n" +
			"    action: transform\n" +
			"    input: ${{ steps.fetch-data.output }}\n" +
			"  - name: notify\n" +
			"    action: slack.send\n" +
			"    channel: '#reports'\n" +
			"    message: ${{ steps.process.output }}\n" +
			"    on_failure: email.alert\n" +
			cb + "\n\n" +
			"### 3. 错误处理\n" +
			"- **重试逻辑**：3 次尝试，指数退避\n" +
			"- **超时**：设置最大执行时间\n" +
			"- **告警**：失败时通知（Slack、邮件）\n" +
			"- **幂等性**：可安全重复执行\n\n" +
			"### 4. 常见自动化场景\n\n" +
			"| 任务 | 工具 | 触发方式 |\n" +
			"|------|------|---------|\n" +
			"| 文件备份 | rsync + cron | 定时 |\n" +
			"| 数据库清理 | SQL 脚本 + cron | 定时 |\n" +
			"| 部署 | GitHub Actions | Git push |\n" +
			"| 日志轮转 | logrotate | 每日 |\n" +
			"| 健康检查 | curl + cron | 每 5 分钟 |\n",
	},
	{
		title: "UI Component Builder", titleZH: "UI 组件构建器",
		desc: "Generate React/Vue components from design descriptions", descZH: "根据设计描述生成 React/Vue 组件",
		category: "Design & Creative", icon: "🎨", tags: "ui,components,react,vue,tailwind", skillType: "ai-powered", featured: 1, likes: 9, views: 54,
		content: "# UI Component Builder\n\n" +
			"Generate production-ready React or Vue components from natural language descriptions.\n\n" +
			"## Component Quality Standards\n\n" +
			"Every generated component must include:\n" +
			"- TypeScript types for all props\n" +
			"- Accessible markup (ARIA labels, keyboard navigation)\n" +
			"- Responsive design (mobile-first)\n" +
			"- Loading and error states\n" +
			"- Clean, composable API\n\n" +
			"## React Template\n\n" +
			cb + "tsx\n" +
			"interface ButtonProps {\n" +
			"  variant?: 'primary' | 'secondary' | 'danger';\n" +
			"  size?: 'sm' | 'md' | 'lg';\n" +
			"  loading?: boolean;\n" +
			"  disabled?: boolean;\n" +
			"  children: React.ReactNode;\n" +
			"  onClick?: () => void;\n" +
			"}\n\n" +
			"export function Button({\n" +
			"  variant = 'primary',\n" +
			"  size = 'md',\n" +
			"  loading = false,\n" +
			"  disabled = false,\n" +
			"  children,\n" +
			"  onClick,\n" +
			"}: ButtonProps) {\n" +
			"  return (\n" +
			"    <button\n" +
			"      className={cn(styles.base, styles[variant], styles[size])}\n" +
			"      disabled={disabled || loading}\n" +
			"      onClick={onClick}\n" +
			"      aria-busy={loading}\n" +
			"    >\n" +
			"      {loading ? <Spinner size={size} /> : children}\n" +
			"    </button>\n" +
			"  );\n" +
			"}\n" +
			cb + "\n\n" +
			"## Design System Tokens\n\n" +
			"Use consistent spacing and color tokens:\n\n" +
			"| Token | Value | Usage |\n" +
			"|-------|-------|-------|\n" +
			"| `space-xs` | 4px | Icon gaps |\n" +
			"| `space-sm` | 8px | Tight padding |\n" +
			"| `space-md` | 16px | Default padding |\n" +
			"| `space-lg` | 24px | Section spacing |\n" +
			"| `radius-sm` | 4px | Buttons, inputs |\n" +
			"| `radius-md` | 8px | Cards |\n" +
			"| `radius-lg` | 16px | Modals |\n\n" +
			"## Accessibility Checklist\n\n" +
			"- [ ] Color contrast ratio ≥ 4.5:1\n" +
			"- [ ] All interactive elements focusable via keyboard\n" +
			"- [ ] ARIA labels on icon-only buttons\n" +
			"- [ ] Form inputs linked to labels\n" +
			"- [ ] Focus trap in modals/dialogs\n",
		contentZH: "# UI 组件构建器\n\n" +
			"根据自然语言描述生成生产级的 React 或 Vue 组件。\n\n" +
			"## 组件质量标准\n\n" +
			"每个生成的组件必须包含：\n" +
			"- 所有 props 的 TypeScript 类型定义\n" +
			"- 无障碍标记（ARIA 标签、键盘导航）\n" +
			"- 响应式设计（移动优先）\n" +
			"- 加载和错误状态\n" +
			"- 干净的、可组合的 API\n\n" +
			"## React 模板\n\n" +
			cb + "tsx\n" +
			"interface ButtonProps {\n" +
			"  variant?: 'primary' | 'secondary' | 'danger';\n" +
			"  size?: 'sm' | 'md' | 'lg';\n" +
			"  loading?: boolean;\n" +
			"  disabled?: boolean;\n" +
			"  children: React.ReactNode;\n" +
			"  onClick?: () => void;\n" +
			"}\n\n" +
			"export function Button({\n" +
			"  variant = 'primary',\n" +
			"  size = 'md',\n" +
			"  loading = false,\n" +
			"  disabled = false,\n" +
			"  children,\n" +
			"  onClick,\n" +
			"}: ButtonProps) {\n" +
			"  return (\n" +
			"    <button\n" +
			"      className={cn(styles.base, styles[variant], styles[size])}\n" +
			"      disabled={disabled || loading}\n" +
			"      onClick={onClick}\n" +
			"      aria-busy={loading}\n" +
			"    >\n" +
			"      {loading ? <Spinner size={size} /> : children}\n" +
			"    </button>\n" +
			"  );\n" +
			"}\n" +
			cb + "\n\n" +
			"## 设计系统令牌\n\n" +
			"使用一致的间距和颜色令牌：\n\n" +
			"| 令牌 | 值 | 用途 |\n" +
			"|------|-----|------|\n" +
			"| `space-xs` | 4px | 图标间距 |\n" +
			"| `space-sm` | 8px | 紧凑内边距 |\n" +
			"| `space-md` | 16px | 默认内边距 |\n" +
			"| `space-lg` | 24px | 区域间距 |\n" +
			"| `radius-sm` | 4px | 按钮、输入框 |\n" +
			"| `radius-md` | 8px | 卡片 |\n" +
			"| `radius-lg` | 16px | 模态框 |\n\n" +
			"## 无障碍清单\n\n" +
			"- [ ] 颜色对比度 ≥ 4.5:1\n" +
			"- [ ] 所有交互元素可通过键盘聚焦\n" +
			"- [ ] 纯图标按钮有 ARIA 标签\n" +
			"- [ ] 表单输入关联 label\n" +
			"- [ ] 模态框中有焦点陷阱\n",
	},
	{
		title: "Database Schema Designer", titleZH: "数据库架构设计师",
		desc: "Design optimal database schemas with relationship mapping", descZH: "设计最优数据库架构，支持关系映射",
		category: "Coding & Development", icon: "🗄️", tags: "database,schema,design,sql,migration", skillType: "general", featured: 0, likes: 5, views: 28,
		content: "# Database Schema Designer\n\n" +
			"Design normalized, performant database schemas from business requirements.\n\n" +
			"## Design Process\n\n" +
			"1. **Identify entities** from requirements (nouns = tables)\n" +
			"2. **Define relationships** (1:1, 1:N, M:N)\n" +
			"3. **Normalize to 3NF** then selectively denormalize for performance\n" +
			"4. **Choose data types** carefully (smallest sufficient type)\n" +
			"5. **Add indexes** for frequent query patterns\n\n" +
			"## Schema Template\n\n" +
			cb + "sql\n" +
			"-- Use UUID or BIGINT for primary keys (not INT)\n" +
			"-- Always include created_at, updated_at\n" +
			"-- Use NOT NULL by default, allow NULL only when needed\n\n" +
			"CREATE TABLE users (\n" +
			"  id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,\n" +
			"  email      VARCHAR(255) NOT NULL UNIQUE,\n" +
			"  name       VARCHAR(100) NOT NULL,\n" +
			"  status     VARCHAR(20) NOT NULL DEFAULT 'active'\n" +
			"             CHECK (status IN ('active','suspended','deleted')),\n" +
			"  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),\n" +
			"  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()\n" +
			");\n\n" +
			"-- M:N relationship via junction table\n" +
			"CREATE TABLE user_roles (\n" +
			"  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,\n" +
			"  role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,\n" +
			"  PRIMARY KEY (user_id, role_id)\n" +
			");\n\n" +
			"-- Index for common queries\n" +
			"CREATE INDEX idx_users_email ON users(email);\n" +
			"CREATE INDEX idx_users_status ON users(status) WHERE status = 'active';\n" +
			cb + "\n\n" +
			"## Index Guidelines\n\n" +
			"| Query Pattern | Index Type |\n" +
			"|---------------|------------|\n" +
			"| WHERE col = val | B-tree (default) |\n" +
			"| WHERE col LIKE 'prefix%' | B-tree |\n" +
			"| Full-text search | GIN / Full-text |\n" +
			"| JSON field queries | GIN |\n" +
			"| Geospatial | GiST / R-tree |\n" +
			"| Composite WHERE a AND b | Composite index (a, b) |\n\n" +
			"## Anti-Patterns to Avoid\n\n" +
			"- ❌ Storing comma-separated values (use junction tables)\n" +
			"- ❌ Using `TEXT` for everything (choose appropriate types)\n" +
			"- ❌ Missing foreign key constraints\n" +
			"- ❌ No indexes on JOIN columns\n" +
			"- ❌ Storing derived data that can be computed\n",
		contentZH: "# 数据库架构设计师\n\n" +
			"根据业务需求设计规范化、高性能的数据库架构。\n\n" +
			"## 设计流程\n\n" +
			"1. **识别实体**（需求中的名词 = 表）\n" +
			"2. **定义关系**（1:1, 1:N, M:N）\n" +
			"3. **规范化到 3NF**，然后针对性能选择性反规范化\n" +
			"4. **选择数据类型**（最小够用类型）\n" +
			"5. **添加索引**（针对常见查询模式）\n\n" +
			"## 架构模板\n\n" +
			cb + "sql\n" +
			"-- 使用 UUID 或 BIGINT 作为主键（而非 INT）\n" +
			"-- 始终包含 created_at, updated_at\n" +
			"-- 默认使用 NOT NULL，仅在需要时允许 NULL\n\n" +
			"CREATE TABLE users (\n" +
			"  id         BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,\n" +
			"  email      VARCHAR(255) NOT NULL UNIQUE,\n" +
			"  name       VARCHAR(100) NOT NULL,\n" +
			"  status     VARCHAR(20) NOT NULL DEFAULT 'active'\n" +
			"             CHECK (status IN ('active','suspended','deleted')),\n" +
			"  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),\n" +
			"  updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()\n" +
			");\n\n" +
			"-- M:N 关系通过关联表实现\n" +
			"CREATE TABLE user_roles (\n" +
			"  user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,\n" +
			"  role_id BIGINT NOT NULL REFERENCES roles(id) ON DELETE CASCADE,\n" +
			"  PRIMARY KEY (user_id, role_id)\n" +
			");\n\n" +
			"-- 常见查询的索引\n" +
			"CREATE INDEX idx_users_email ON users(email);\n" +
			"CREATE INDEX idx_users_status ON users(status) WHERE status = 'active';\n" +
			cb + "\n\n" +
			"## 索引指南\n\n" +
			"| 查询模式 | 索引类型 |\n" +
			"|----------|----------|\n" +
			"| WHERE col = val | B-tree（默认） |\n" +
			"| WHERE col LIKE 'prefix%' | B-tree |\n" +
			"| 全文搜索 | GIN / 全文索引 |\n" +
			"| JSON 字段查询 | GIN |\n" +
			"| 地理空间 | GiST / R-tree |\n" +
			"| 复合 WHERE a AND b | 复合索引 (a, b) |\n\n" +
			"## 反模式\n\n" +
			"- ❌ 存储逗号分隔的值（应使用关联表）\n" +
			"- ❌ 所有字段都用 TEXT（选择合适的类型）\n" +
			"- ❌ 缺少外键约束\n" +
			"- ❌ JOIN 列没有索引\n" +
			"- ❌ 存储可计算的派生数据\n",
	},
	{
		title: "SEO Optimizer", titleZH: "SEO 优化器",
		desc: "Analyze and optimize content for search engines", descZH: "分析和优化搜索引擎内容",
		category: "Writing & Content", icon: "🔎", tags: "seo,optimization,content,marketing", skillType: "ai-powered", featured: 0, likes: 4, views: 24,
		content: "# SEO Optimizer\n\n" +
			"Analyze web pages and content for search engine optimization opportunities.\n\n" +
			"## On-Page SEO Checklist\n\n" +
			"### Title Tag\n" +
			"- [ ] Contains primary keyword\n" +
			"- [ ] Under 60 characters\n" +
			"- [ ] Unique across the site\n" +
			"- [ ] Compelling (includes number, power word, or question)\n\n" +
			"### Meta Description\n" +
			"- [ ] 120-160 characters\n" +
			"- [ ] Contains primary keyword\n" +
			"- [ ] Includes a call-to-action\n" +
			"- [ ] Accurately describes page content\n\n" +
			"### Headings\n" +
			"- [ ] Single H1 tag containing primary keyword\n" +
			"- [ ] H2s contain secondary keywords\n" +
			"- [ ] Logical heading hierarchy (H1 → H2 → H3)\n\n" +
			"### Content\n" +
			"- [ ] Minimum 300 words (1500+ for competitive keywords)\n" +
			"- [ ] Primary keyword density: 1-2%\n" +
			"- [ ] LSI keywords used naturally\n" +
			"- [ ] Short paragraphs (3-4 sentences)\n" +
			"- [ ] Includes lists, tables, or structured data\n\n" +
			"### Technical\n" +
			"- [ ] Page loads in < 3 seconds\n" +
			"- [ ] Mobile-friendly (responsive design)\n" +
			"- [ ] Images optimized (WebP, lazy loading, alt text)\n" +
			"- [ ] Internal links to related pages\n" +
			"- [ ] External links to authority sources\n" +
			"- [ ] Canonical URL set\n" +
			"- [ ] Schema markup (JSON-LD) for rich snippets\n\n" +
			"## Schema Markup Example\n\n" +
			cb + "json\n" +
			"{\n" +
			"  \"@context\": \"https://schema.org\",\n" +
			"  \"@type\": \"Article\",\n" +
			"  \"headline\": \"Your Article Title\",\n" +
			"  \"author\": { \"@type\": \"Person\", \"name\": \"Author\" },\n" +
			"  \"datePublished\": \"2026-01-01\",\n" +
			"  \"image\": \"https://example.com/image.jpg\"\n" +
			"}\n" +
			cb + "\n",
		contentZH: "# SEO 优化器\n\n" +
			"分析网页和内容，提供搜索引擎优化建议。\n\n" +
			"## 页面 SEO 清单\n\n" +
			"### 标题标签\n" +
			"- [ ] 包含主关键词\n" +
			"- [ ] 60 字符以内\n" +
			"- [ ] 全站唯一\n" +
			"- [ ] 有吸引力（包含数字、有力词汇或疑问句）\n\n" +
			"### Meta 描述\n" +
			"- [ ] 120-160 字符\n" +
			"- [ ] 包含主关键词和行动号召\n" +
			"- [ ] 准确描述页面内容\n\n" +
			"### 标题层级\n" +
			"- [ ] 单个 H1 标签包含主关键词\n" +
			"- [ ] H2 包含次要关键词\n" +
			"- [ ] 逻辑标题层级（H1 → H2 → H3）\n\n" +
			"### 内容\n" +
			"- [ ] 至少 300 字（竞争性关键词 1500+）\n" +
			"- [ ] 关键词密度 1-2%\n" +
			"- [ ] 自然使用相关关键词\n" +
			"- [ ] 短段落（3-4 句）\n" +
			"- [ ] 包含列表、表格或结构化数据\n\n" +
			"### 技术\n" +
			"- [ ] 页面加载 < 3 秒\n" +
			"- [ ] 移动端适配（响应式设计）\n" +
			"- [ ] 图片优化（WebP、懒加载、alt 文本）\n" +
			"- [ ] 内链到相关页面\n" +
			"- [ ] 外链到权威来源\n" +
			"- [ ] 设置规范 URL\n" +
			"- [ ] 结构化数据标记（JSON-LD）用于富摘要\n\n" +
			"## Schema 标记示例\n\n" +
			cb + "json\n" +
			"{\n" +
			"  \"@context\": \"https://schema.org\",\n" +
			"  \"@type\": \"Article\",\n" +
			"  \"headline\": \"Your Article Title\",\n" +
			"  \"author\": { \"@type\": \"Person\", \"name\": \"Author\" },\n" +
			"  \"datePublished\": \"2026-01-01\",\n" +
			"  \"image\": \"https://example.com/image.jpg\"\n" +
			"}\n" +
			cb + "\n",
	},
	{
		title: "DevOps Pipeline Setup", titleZH: "DevOps 流水线配置",
		desc: "Configure CI/CD pipelines with best practices", descZH: "配置 CI/CD 流水线，包含最佳实践",
		category: "Automation", icon: "⚙️", tags: "devops,cicd,pipeline,github-actions,docker", skillType: "general", featured: 0, likes: 5, views: 32,
		content: "# DevOps Pipeline Setup\n\n" +
			"Configure production-grade CI/CD pipelines for any project.\n\n" +
			"## GitHub Actions Template\n\n" +
			cb + "yaml\n" +
			"name: CI/CD\n" +
			"on:\n" +
			"  push:\n" +
			"    branches: [main]\n" +
			"  pull_request:\n" +
			"    branches: [main]\n\n" +
			"jobs:\n" +
			"  test:\n" +
			"    runs-on: ubuntu-latest\n" +
			"    steps:\n" +
			"      - uses: actions/checkout@v4\n" +
			"      - uses: actions/setup-node@v4\n" +
			"        with: { node-version: 20, cache: npm }\n" +
			"      - run: npm ci\n" +
			"      - run: npm run lint\n" +
			"      - run: npm test -- --coverage\n" +
			"      - uses: actions/upload-artifact@v4\n" +
			"        with:\n" +
			"          name: coverage\n" +
			"          path: coverage/\n\n" +
			"  deploy:\n" +
			"    needs: test\n" +
			"    if: github.ref == 'refs/heads/main'\n" +
			"    runs-on: ubuntu-latest\n" +
			"    steps:\n" +
			"      - uses: actions/checkout@v4\n" +
			"      - run: docker build -t app:${{ github.sha }} .\n" +
			"      - run: docker push registry/app:${{ github.sha }}\n" +
			"      - run: kubectl set image deploy/app app=registry/app:${{ github.sha }}\n" +
			cb + "\n\n" +
			"## Pipeline Stages\n\n" +
			"| Stage | Purpose | Fail = Block? |\n" +
			"|-------|---------|---------------|\n" +
			"| Lint | Code style | Yes |\n" +
			"| Test | Unit + integration tests | Yes |\n" +
			"| Build | Compile / bundle | Yes |\n" +
			"| Security | Dependency audit (npm audit) | Warning |\n" +
			"| Deploy staging | Preview environment | Yes |\n" +
			"| Deploy prod | Production release | Yes |\n\n" +
			"## Best Practices\n\n" +
			"- Cache dependencies (`actions/cache`) to speed up builds\n" +
			"- Use branch protection rules: require passing CI before merge\n" +
			"- Pin action versions to SHA for security\n" +
			"- Keep secrets in GitHub Secrets, never in code\n" +
			"- Use matrix builds for multiple OS/language versions\n" +
			"- Set timeouts to prevent hung jobs\n",
		contentZH: "# DevOps 流水线配置\n\n" +
			"为项目配置生产级 CI/CD 流水线。\n\n" +
			"## GitHub Actions 模板\n\n" +
			cb + "yaml\n" +
			"name: CI/CD\n" +
			"on:\n" +
			"  push:\n" +
			"    branches: [main]\n" +
			"  pull_request:\n" +
			"    branches: [main]\n\n" +
			"jobs:\n" +
			"  test:\n" +
			"    runs-on: ubuntu-latest\n" +
			"    steps:\n" +
			"      - uses: actions/checkout@v4\n" +
			"      - uses: actions/setup-node@v4\n" +
			"        with: { node-version: 20, cache: npm }\n" +
			"      - run: npm ci\n" +
			"      - run: npm run lint\n" +
			"      - run: npm test -- --coverage\n" +
			"      - uses: actions/upload-artifact@v4\n" +
			"        with:\n" +
			"          name: coverage\n" +
			"          path: coverage/\n\n" +
			"  deploy:\n" +
			"    needs: test\n" +
			"    if: github.ref == 'refs/heads/main'\n" +
			"    runs-on: ubuntu-latest\n" +
			"    steps:\n" +
			"      - uses: actions/checkout@v4\n" +
			"      - run: docker build -t app:${{ github.sha }} .\n" +
			"      - run: docker push registry/app:${{ github.sha }}\n" +
			"      - run: kubectl set image deploy/app app=registry/app:${{ github.sha }}\n" +
			cb + "\n\n" +
			"## 流水线阶段\n\n" +
			"| 阶段 | 目的 | 失败是否阻断？ |\n" +
			"|------|------|---------------|\n" +
			"| Lint | 代码风格 | 是 |\n" +
			"| 测试 | 单元 + 集成测试 | 是 |\n" +
			"| 构建 | 编译/打包 | 是 |\n" +
			"| 安全 | 依赖审计 | 警告 |\n" +
			"| 部署预发 | 预览环境 | 是 |\n" +
			"| 部署生产 | 生产发布 | 是 |\n\n" +
			"## 最佳实践\n\n" +
			"- 缓存依赖以加速构建\n" +
			"- 使用分支保护规则：合并前要求 CI 通过\n" +
			"- 固定 Action 版本到 SHA\n" +
			"- 密钥存储在 GitHub Secrets 中\n" +
			"- 使用矩阵构建覆盖多个 OS/语言版本\n" +
			"- 设置超时防止作业挂起\n",
	},

	// === New Skills ===

	{
		title: "Git Commit Message Generator", titleZH: "Git 提交信息生成器",
		desc: "Generate clear, conventional commit messages from code diffs", descZH: "根据代码差异生成清晰规范的提交信息",
		category: "Coding & Development", icon: "📝", tags: "git,commit,conventional-commits,version-control", skillType: "ai-powered", featured: 1, likes: 9, views: 58,
		content: "# Git Commit Message Generator\n\n" +
			"Generate descriptive commit messages following the Conventional Commits specification.\n\n" +
			"## Format\n\n" +
			cb + "\n" +
			"<type>(<scope>): <subject>\n\n" +
			"<body>\n\n" +
			"<footer>\n" +
			cb + "\n\n" +
			"## Types\n\n" +
			"| Type | When to Use |\n" +
			"|------|-------------|\n" +
			"| `feat` | New feature for the user |\n" +
			"| `fix` | Bug fix |\n" +
			"| `docs` | Documentation only |\n" +
			"| `style` | Formatting, no logic change |\n" +
			"| `refactor` | Code restructure, no behavior change |\n" +
			"| `perf` | Performance improvement |\n" +
			"| `test` | Adding or fixing tests |\n" +
			"| `chore` | Build, CI, dependencies |\n\n" +
			"## Rules\n\n" +
			"1. **Subject**: Imperative mood, < 50 chars, no period\n" +
			"   - ✅ `feat(auth): add JWT token refresh`\n" +
			"   - ❌ `feat(auth): Added JWT token refresh.`\n\n" +
			"2. **Scope**: The module/area affected (optional but recommended)\n" +
			"   - Examples: `auth`, `api`, `ui`, `db`, `config`\n\n" +
			"3. **Body**: Explain *why*, not *what* (the diff shows what)\n" +
			"   - Wrap at 72 characters\n" +
			"   - Separate from subject with blank line\n\n" +
			"4. **Footer**: Reference issues, breaking changes\n" +
			"   - `Closes #123`\n" +
			"   - `BREAKING CHANGE: API response format changed`\n\n" +
			"## Examples\n\n" +
			cb + "\n" +
			"feat(cart): add quantity validation on checkout\n\n" +
			"Prevent orders with quantity > stock from being placed.\n" +
			"Previously this was only validated client-side.\n\n" +
			"Closes #456\n" +
			cb + "\n\n" +
			cb + "\n" +
			"fix(api): handle timeout in payment webhook\n\n" +
			"Payment provider occasionally sends delayed webhooks\n" +
			"that exceeded our 5s timeout. Increased to 30s and\n" +
			"added retry logic.\n" +
			cb + "\n",
		contentZH: "# Git 提交信息生成器\n\n" +
			"按照 Conventional Commits 规范生成描述性的提交信息。\n\n" +
			"## 格式\n\n" +
			cb + "\n" +
			"<类型>(<范围>): <主题>\n\n" +
			"<正文>\n\n" +
			"<页脚>\n" +
			cb + "\n\n" +
			"## 类型\n\n" +
			"| 类型 | 使用场景 |\n" +
			"|------|----------|\n" +
			"| `feat` | 新功能 |\n" +
			"| `fix` | 修复 bug |\n" +
			"| `docs` | 仅文档 |\n" +
			"| `refactor` | 重构，不改变行为 |\n" +
			"| `perf` | 性能优化 |\n" +
			"| `test` | 添加或修复测试 |\n" +
			"| `chore` | 构建、CI、依赖 |\n\n" +
			"## 规则\n\n" +
			"1. **主题**：使用祈使句，< 50 字符，不加句号\n" +
			"   - ✅ `feat(auth): add JWT token refresh`\n" +
			"   - ❌ `feat(auth): Added JWT token refresh.`\n\n" +
			"2. **范围**：受影响的模块（可选但推荐）\n" +
			"   - 例如：`auth`, `api`, `ui`, `db`, `config`\n\n" +
			"3. **正文**：解释*为什么*，而不是*做了什么*（diff 已经展示了做了什么）\n" +
			"   - 每行不超过 72 字符\n" +
			"   - 与主题之间空一行\n\n" +
			"4. **页脚**：引用 issue、标注破坏性变更\n" +
			"   - `Closes #123`\n" +
			"   - `BREAKING CHANGE: API 响应格式变更`\n\n" +
			"## 示例\n\n" +
			cb + "\n" +
			"feat(cart): add quantity validation on checkout\n\n" +
			"Prevent orders with quantity > stock from being placed.\n" +
			"Previously this was only validated client-side.\n\n" +
			"Closes #456\n" +
			cb + "\n\n" +
			cb + "\n" +
			"fix(api): handle timeout in payment webhook\n\n" +
			"Payment provider occasionally sends delayed webhooks\n" +
			"that exceeded our 5s timeout. Increased to 30s and\n" +
			"added retry logic.\n" +
			cb + "\n",
	},
	{
		title: "README Generator", titleZH: "README 生成器",
		desc: "Generate comprehensive README files for any project", descZH: "为任何项目生成全面的 README 文件",
		category: "Writing & Content", icon: "📄", tags: "readme,documentation,markdown,open-source", skillType: "ai-powered", featured: 0, likes: 6, views: 39,
		content: "# README Generator\n\n" +
			"Generate a comprehensive, well-structured README.md for any project.\n\n" +
			"## Template\n\n" +
			cb + "markdown\n" +
			"# Project Name\n\n" +
			"One-line description of what this project does.\n\n" +
			"![Build Status](badge-url) ![License](badge-url)\n\n" +
			"## Features\n\n" +
			"- Feature 1: Brief description\n" +
			"- Feature 2: Brief description\n\n" +
			"## Quick Start\n\n" +
			"### Prerequisites\n" +
			"- Node.js >= 18\n" +
			"- PostgreSQL >= 14\n\n" +
			"### Installation\n" +
			"git clone https://github.com/user/repo.git\n" +
			"cd repo\n" +
			"npm install\n" +
			"cp .env.example .env\n" +
			"npm run dev\n\n" +
			"## Usage\n\n" +
			"Brief usage example with code.\n\n" +
			"## API Reference\n\n" +
			"| Endpoint | Method | Description |\n" +
			"|----------|--------|-------------|\n" +
			"| /api/users | GET | List users |\n\n" +
			"## Contributing\n\n" +
			"1. Fork the repository\n" +
			"2. Create your feature branch\n" +
			"3. Commit your changes\n" +
			"4. Push to the branch\n" +
			"5. Open a Pull Request\n\n" +
			"## License\n\n" +
			"MIT License - see LICENSE file\n" +
			cb + "\n\n" +
			"## Section Guidelines\n\n" +
			"| Section | Required? | Notes |\n" +
			"|---------|-----------|-------|\n" +
			"| Title + description | Yes | First thing people see |\n" +
			"| Quick Start | Yes | Get running in < 5 min |\n" +
			"| Features | Yes | Why should I use this? |\n" +
			"| Usage/Examples | Yes | Show, don't just tell |\n" +
			"| API Reference | If applicable | Document all public APIs |\n" +
			"| Contributing | For open source | Lower the barrier |\n" +
			"| License | Yes | Legal clarity |\n",
		contentZH: "# README 生成器\n\n" +
			"为任何项目生成结构完整的 README.md 文件。\n\n" +
			"## 模板\n\n" +
			cb + "markdown\n" +
			"# Project Name\n\n" +
			"One-line description of what this project does.\n\n" +
			"![Build Status](badge-url) ![License](badge-url)\n\n" +
			"## Features\n\n" +
			"- Feature 1: Brief description\n" +
			"- Feature 2: Brief description\n\n" +
			"## Quick Start\n\n" +
			"### Prerequisites\n" +
			"- Node.js >= 18\n" +
			"- PostgreSQL >= 14\n\n" +
			"### Installation\n" +
			"git clone https://github.com/user/repo.git\n" +
			"cd repo\n" +
			"npm install\n" +
			"cp .env.example .env\n" +
			"npm run dev\n\n" +
			"## Usage\n\n" +
			"Brief usage example with code.\n\n" +
			"## API Reference\n\n" +
			"| Endpoint | Method | Description |\n" +
			"|----------|--------|-------------|\n" +
			"| /api/users | GET | List users |\n\n" +
			"## Contributing\n\n" +
			"1. Fork the repository\n" +
			"2. Create your feature branch\n" +
			"3. Commit your changes\n" +
			"4. Push to the branch\n" +
			"5. Open a Pull Request\n\n" +
			"## License\n\n" +
			"MIT License - see LICENSE file\n" +
			cb + "\n\n" +
			"## 章节指南\n\n" +
			"| 章节 | 是否必需 | 说明 |\n" +
			"|------|----------|------|\n" +
			"| 标题 + 描述 | 是 | 用户第一眼看到的内容 |\n" +
			"| 快速开始 | 是 | 5 分钟内跑起来 |\n" +
			"| 功能特性 | 是 | 为什么要用这个项目？ |\n" +
			"| 使用示例 | 是 | 展示而不只是告知 |\n" +
			"| API 文档 | 如适用 | 文档化所有公开 API |\n" +
			"| 贡献指南 | 开源项目 | 降低参与门槛 |\n" +
			"| 许可证 | 是 | 法律明确性 |\n",
	},
	{
		title: "Regex Pattern Builder", titleZH: "正则表达式构建器",
		desc: "Build and explain complex regular expressions step by step", descZH: "逐步构建和解释复杂的正则表达式",
		category: "Coding & Development", icon: "🔤", tags: "regex,pattern-matching,validation,text-processing", skillType: "ai-powered", featured: 0, likes: 5, views: 34,
		content: "# Regex Pattern Builder\n\n" +
			"Build, test, and explain regular expressions step by step.\n\n" +
			"## Common Patterns\n\n" +
			"| Need | Pattern | Notes |\n" +
			"|------|---------|-------|\n" +
			"| Email | `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$` | Basic validation |\n" +
			"| URL | `https?://[^\\s/$.?#].[^\\s]*` | HTTP/HTTPS URLs |\n" +
			"| Phone (US) | `\\(?\\d{3}\\)?[-.\\s]?\\d{3}[-.\\s]?\\d{4}` | Multiple formats |\n" +
			"| IPv4 | `\\b\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\b` | Basic |\n" +
			"| Date (ISO) | `\\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\\d|3[01])` | YYYY-MM-DD |\n" +
			"| Hex Color | `#(?:[0-9a-fA-F]{3}){1,2}` | #RGB or #RRGGBB |\n" +
			"| Password | `^(?=.*[a-z])(?=.*[A-Z])(?=.*\\d)(?=.*[@$!%]).{8,}$` | Strong |\n\n" +
			"## Building Blocks\n\n" +
			"| Symbol | Meaning | Example |\n" +
			"|--------|---------|--------|\n" +
			"| `.` | Any character | `a.c` → abc, a1c |\n" +
			"| `*` | 0 or more | `ab*c` → ac, abc, abbc |\n" +
			"| `+` | 1 or more | `ab+c` → abc, abbc |\n" +
			"| `?` | 0 or 1 | `colou?r` → color, colour |\n" +
			"| `{n,m}` | n to m times | `a{2,4}` → aa, aaa, aaaa |\n" +
			"| `()` | Capture group | `(ab)+` → ab, abab |\n" +
			"| `(?:)` | Non-capturing group | `(?:ab)+` → ab, abab |\n" +
			"| `(?=)` | Lookahead | `foo(?=bar)` → foo in foobar |\n" +
			"| `(?<=)` | Lookbehind | `(?<=@)\\w+` → domain in @domain |\n" +
			"| `[^]` | Negated set | `[^0-9]` → not a digit |\n\n" +
			"## Output Format\n\n" +
			"When building a regex, always provide:\n" +
			"1. The pattern\n" +
			"2. Step-by-step explanation of each part\n" +
			"3. Test cases (matches and non-matches)\n" +
			"4. Performance notes (avoid catastrophic backtracking)\n",
		contentZH: "# 正则表达式构建器\n\n" +
			"逐步构建、测试和解释正则表达式。\n\n" +
			"## 常用模式\n\n" +
			"| 需求 | 模式 |\n" +
			"|------|------|\n" +
			"| 邮箱 | `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,}$` |\n" +
			"| URL | `https?://[^\\\\s/$.?#].[^\\\\s]*` |\n" +
			"| 手机号（中国） | `^1[3-9]\\\\d{9}$` |\n" +
			"| 日期 | `\\\\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\\\\d|3[01])` |\n\n" +
			"## 构建块\n\n" +
			"| 符号 | 含义 | 示例 |\n" +
			"|------|------|------|\n" +
			"| `.` | 任意字符 | `a.c` → abc, a1c |\n" +
			"| `*` | 0 个或多个 | `ab*c` → ac, abc, abbc |\n" +
			"| `+` | 1 个或多个 | `ab+c` → abc, abbc |\n" +
			"| `?` | 0 个或 1 个 | `colou?r` → color, colour |\n" +
			"| `{n,m}` | n 到 m 次 | `a{2,4}` → aa, aaa, aaaa |\n" +
			"| `()` | 捕获组 | `(ab)+` → ab, abab |\n" +
			"| `(?:)` | 非捕获组 | `(?:ab)+` → ab, abab |\n" +
			"| `(?=)` | 前瞻断言 | `foo(?=bar)` → foobar 中的 foo |\n" +
			"| `(?<=)` | 后瞻断言 | `(?<=@)\\\\w+` → @domain 中的 domain |\n" +
			"| `[^]` | 否定集合 | `[^0-9]` → 非数字 |\n\n" +
			"## 输出格式\n\n" +
			"构建正则时，始终提供：\n" +
			"1. 正则模式\n" +
			"2. 每部分的逐步解释\n" +
			"3. 测试用例（匹配和不匹配的）\n" +
			"4. 性能说明（避免灾难性回溯）\n",
	},
	{
		title: "SQL Query Optimizer", titleZH: "SQL 查询优化器",
		desc: "Analyze and optimize SQL queries for better performance", descZH: "分析和优化 SQL 查询以提升性能",
		category: "Data & Analytics", icon: "⚡", tags: "sql,performance,optimization,database,query", skillType: "ai-powered", featured: 1, likes: 8, views: 50,
		content: "# SQL Query Optimizer\n\n" +
			"Analyze SQL queries and suggest performance optimizations.\n\n" +
			"## Optimization Checklist\n\n" +
			"### 1. Check EXPLAIN Plan\n\n" +
			cb + "sql\n" +
			"EXPLAIN ANALYZE SELECT * FROM orders\n" +
			"WHERE user_id = 123 AND status = 'active'\n" +
			"ORDER BY created_at DESC LIMIT 20;\n" +
			cb + "\n\n" +
			"Look for:\n" +
			"- **Seq Scan** on large tables → Add an index\n" +
			"- **Nested Loop** with large tables → Consider Hash Join\n" +
			"- **Sort** with high cost → Add index matching ORDER BY\n\n" +
			"### 2. Common Optimizations\n\n" +
			"| Problem | Solution |\n" +
			"|---------|----------|\n" +
			"| `SELECT *` | Select only needed columns |\n" +
			"| No WHERE index | Add composite index matching WHERE clause |\n" +
			"| `LIKE '%term%'` | Use full-text search or trigram index |\n" +
			"| Subquery in WHERE | Rewrite as JOIN |\n" +
			"| `DISTINCT` on large set | Fix the JOIN causing duplicates |\n" +
			"| `COUNT(*)` on full table | Use approximate count or materialized view |\n" +
			"| N+1 in app code | Use JOIN or batch query |\n\n" +
			"### 3. Index Strategy\n\n" +
			cb + "sql\n" +
			"-- Composite index: put equality columns first, range columns last\n" +
			"-- For: WHERE status = 'active' AND created_at > '2026-01-01'\n" +
			"CREATE INDEX idx_orders_status_date ON orders(status, created_at);\n\n" +
			"-- Covering index: include all SELECT columns to avoid table lookup\n" +
			"CREATE INDEX idx_orders_cover ON orders(user_id, status)\n" +
			"  INCLUDE (total, created_at);\n\n" +
			"-- Partial index: only index rows that matter\n" +
			"CREATE INDEX idx_orders_active ON orders(user_id)\n" +
			"  WHERE status = 'active';\n" +
			cb + "\n\n" +
			"### 4. Query Rewriting\n\n" +
			cb + "sql\n" +
			"-- Before: Slow subquery\n" +
			"SELECT * FROM orders\n" +
			"WHERE user_id IN (SELECT id FROM users WHERE country = 'US');\n\n" +
			"-- After: Faster JOIN\n" +
			"SELECT o.* FROM orders o\n" +
			"JOIN users u ON o.user_id = u.id\n" +
			"WHERE u.country = 'US';\n" +
			cb + "\n",
		contentZH: "# SQL 查询优化器\n\n" +
			"分析 SQL 查询并提供性能优化建议。\n\n" +
			"## 优化清单\n\n" +
			"### 1. 检查执行计划\n\n" +
			cb + "sql\n" +
			"EXPLAIN ANALYZE SELECT * FROM orders\n" +
			"WHERE user_id = 123 AND status = 'active'\n" +
			"ORDER BY created_at DESC LIMIT 20;\n" +
			cb + "\n\n" +
			"注意：\n" +
			"- **全表扫描** → 添加索引\n" +
			"- **嵌套循环** → 考虑 Hash Join\n" +
			"- **排序成本高** → 添加匹配 ORDER BY 的索引\n\n" +
			"### 2. 常见优化\n\n" +
			"| 问题 | 解决方案 |\n" +
			"|------|----------|\n" +
			"| `SELECT *` | 只查询需要的列 |\n" +
			"| WHERE 无索引 | 添加复合索引 |\n" +
			"| `LIKE '%term%'` | 使用全文搜索 |\n" +
			"| WHERE 中的子查询 | 改写为 JOIN |\n" +
			"| 大表 `COUNT(*)` | 使用近似计数或物化视图 |\n" +
			"| N+1 查询 | 使用 JOIN 或批量查询 |\n\n" +
			"### 3. 索引策略\n\n" +
			cb + "sql\n" +
			"-- 复合索引：等值列在前，范围列在后\n" +
			"-- 适用于: WHERE status = 'active' AND created_at > '2026-01-01'\n" +
			"CREATE INDEX idx_orders_status_date ON orders(status, created_at);\n\n" +
			"-- 覆盖索引：包含所有 SELECT 列以避免回表\n" +
			"CREATE INDEX idx_orders_cover ON orders(user_id, status)\n" +
			"  INCLUDE (total, created_at);\n\n" +
			"-- 部分索引：只索引相关的行\n" +
			"CREATE INDEX idx_orders_active ON orders(user_id)\n" +
			"  WHERE status = 'active';\n" +
			cb + "\n\n" +
			"### 4. 查询改写\n\n" +
			cb + "sql\n" +
			"-- 改写前：慢子查询\n" +
			"SELECT * FROM orders\n" +
			"WHERE user_id IN (SELECT id FROM users WHERE country = 'US');\n\n" +
			"-- 改写后：更快的 JOIN\n" +
			"SELECT o.* FROM orders o\n" +
			"JOIN users u ON o.user_id = u.id\n" +
			"WHERE u.country = 'US';\n" +
			cb + "\n",
	},
	{
		title: "Docker Compose Generator", titleZH: "Docker Compose 生成器",
		desc: "Generate Docker Compose configurations for development and production", descZH: "生成开发和生产环境的 Docker Compose 配置",
		category: "Automation", icon: "🐳", tags: "docker,compose,containers,devops,infrastructure", skillType: "general", featured: 0, likes: 6, views: 37,
		content: "# Docker Compose Generator\n\n" +
			"Generate Docker Compose configurations for common tech stacks.\n\n" +
			"## Full-Stack Template\n\n" +
			cb + "yaml\n" +
			"services:\n" +
			"  app:\n" +
			"    build: .\n" +
			"    ports: ['3000:3000']\n" +
			"    environment:\n" +
			"      DATABASE_URL: postgres://user:pass@db:5432/app\n" +
			"      REDIS_URL: redis://cache:6379\n" +
			"    depends_on:\n" +
			"      db: { condition: service_healthy }\n" +
			"      cache: { condition: service_started }\n" +
			"    volumes: ['./src:/app/src']  # Dev hot-reload\n\n" +
			"  db:\n" +
			"    image: postgres:16-alpine\n" +
			"    environment:\n" +
			"      POSTGRES_USER: user\n" +
			"      POSTGRES_PASSWORD: pass\n" +
			"      POSTGRES_DB: app\n" +
			"    volumes: ['pgdata:/var/lib/postgresql/data']\n" +
			"    healthcheck:\n" +
			"      test: pg_isready -U user\n" +
			"      interval: 5s\n" +
			"      retries: 5\n\n" +
			"  cache:\n" +
			"    image: redis:7-alpine\n" +
			"    volumes: ['redisdata:/data']\n\n" +
			"volumes:\n" +
			"  pgdata:\n" +
			"  redisdata:\n" +
			cb + "\n\n" +
			"## Best Practices\n\n" +
			"- Use `depends_on` with `condition: service_healthy` for startup order\n" +
			"- Use Alpine-based images for smaller footprint\n" +
			"- Named volumes for data persistence\n" +
			"- `.env` file for secrets (never commit to git)\n" +
			"- Use multi-stage Dockerfile for production builds\n" +
			"- Set resource limits for production\n\n" +
			"## Common Services\n\n" +
			"| Service | Image | Default Port |\n" +
			"|---------|-------|--------------|\n" +
			"| PostgreSQL | postgres:16-alpine | 5432 |\n" +
			"| MySQL | mysql:8 | 3306 |\n" +
			"| Redis | redis:7-alpine | 6379 |\n" +
			"| MongoDB | mongo:7 | 27017 |\n" +
			"| Elasticsearch | elasticsearch:8 | 9200 |\n" +
			"| RabbitMQ | rabbitmq:3-management | 5672/15672 |\n" +
			"| MinIO | minio/minio | 9000/9001 |\n" +
			"| Mailpit | axllent/mailpit | 1025/8025 |\n",
		contentZH: "# Docker Compose 生成器\n\n" +
			"为常见技术栈生成 Docker Compose 配置。\n\n" +
			"## 全栈模板\n\n" +
			cb + "yaml\n" +
			"services:\n" +
			"  app:\n" +
			"    build: .\n" +
			"    ports: ['3000:3000']\n" +
			"    environment:\n" +
			"      DATABASE_URL: postgres://user:pass@db:5432/app\n" +
			"      REDIS_URL: redis://cache:6379\n" +
			"    depends_on:\n" +
			"      db: { condition: service_healthy }\n" +
			"      cache: { condition: service_started }\n" +
			"    volumes: ['./src:/app/src']  # Dev hot-reload\n\n" +
			"  db:\n" +
			"    image: postgres:16-alpine\n" +
			"    environment:\n" +
			"      POSTGRES_USER: user\n" +
			"      POSTGRES_PASSWORD: pass\n" +
			"      POSTGRES_DB: app\n" +
			"    volumes: ['pgdata:/var/lib/postgresql/data']\n" +
			"    healthcheck:\n" +
			"      test: pg_isready -U user\n" +
			"      interval: 5s\n" +
			"      retries: 5\n\n" +
			"  cache:\n" +
			"    image: redis:7-alpine\n" +
			"    volumes: ['redisdata:/data']\n\n" +
			"volumes:\n" +
			"  pgdata:\n" +
			"  redisdata:\n" +
			cb + "\n\n" +
			"## 最佳实践\n\n" +
			"- 使用 `depends_on` + `condition: service_healthy` 控制启动顺序\n" +
			"- 使用 Alpine 镜像减小体积\n" +
			"- 使用命名卷持久化数据\n" +
			"- 密钥放在 `.env` 文件中（不提交到 git）\n" +
			"- 生产环境使用多阶段 Dockerfile\n" +
			"- 生产环境设置资源限制\n\n" +
			"## 常用服务\n\n" +
			"| 服务 | 镜像 | 默认端口 |\n" +
			"|------|------|---------|\n" +
			"| PostgreSQL | postgres:16-alpine | 5432 |\n" +
			"| MySQL | mysql:8 | 3306 |\n" +
			"| Redis | redis:7-alpine | 6379 |\n" +
			"| MongoDB | mongo:7 | 27017 |\n" +
			"| RabbitMQ | rabbitmq:3-management | 5672 |\n",
	},
	{
		title: "Prompt Engineering Guide", titleZH: "Prompt 工程指南",
		desc: "Master techniques for writing effective AI prompts", descZH: "掌握编写高效 AI 提示词的技巧",
		category: "Data & Analytics", icon: "🧠", tags: "prompt,ai,llm,chatgpt,claude,engineering", skillType: "ai-powered", featured: 1, likes: 10, views: 65,
		content: "# Prompt Engineering Guide\n\n" +
			"Techniques for writing effective prompts that get accurate, useful AI responses.\n\n" +
			"## Core Techniques\n\n" +
			"### 1. Be Specific\n\n" +
			"| Vague ❌ | Specific ✅ |\n" +
			"|----------|------------|\n" +
			"| \"Write a function\" | \"Write a TypeScript function that validates email addresses using regex, returns boolean, handles edge cases like '+' and dots\" |\n" +
			"| \"Make it better\" | \"Improve the error handling: add try-catch, return typed error objects, log to stderr\" |\n" +
			"| \"Explain this code\" | \"Explain the time complexity and memory usage of this sorting algorithm\" |\n\n" +
			"### 2. Provide Context\n\n" +
			"Include:\n" +
			"- **Role**: \"You are a senior backend engineer reviewing Go code\"\n" +
			"- **Constraints**: \"Must be compatible with Node.js 18, no external dependencies\"\n" +
			"- **Format**: \"Output as a markdown table with columns: issue, severity, fix\"\n" +
			"- **Examples**: Show input/output pairs for the expected behavior\n\n" +
			"### 3. Chain of Thought\n\n" +
			"For complex problems, ask the AI to think step by step:\n\n" +
			"\"Analyze this database schema for performance issues. For each issue:\n" +
			"1. Identify the problem\n" +
			"2. Explain why it's a problem\n" +
			"3. Suggest a specific fix with SQL\n" +
			"4. Estimate the performance impact\"\n\n" +
			"### 4. Few-Shot Examples\n\n" +
			"Provide 2-3 examples of desired output:\n\n" +
			"\"Convert requirements to user stories:\n\n" +
			"Requirement: Users can reset passwords\n" +
			"Story: As a user, I want to reset my password via email so that I can regain access to my account.\n\n" +
			"Requirement: Admin can ban users\n" +
			"Story: As an admin, I want to ban users by ID so that I can enforce community guidelines.\n\n" +
			"Now convert: Users can export data as CSV\"\n\n" +
			"### 5. Iterative Refinement\n\n" +
			"- Start with a broad prompt, then narrow down\n" +
			"- Use \"Keep X, but change Y\" for incremental adjustments\n" +
			"- Ask \"What assumptions did you make?\" to surface gaps\n\n" +
			"## Anti-Patterns\n\n" +
			"- ❌ Too many instructions at once (break into steps)\n" +
			"- ❌ Ambiguous pronouns (\"make it work with that\")\n" +
			"- ❌ No success criteria (\"make it good\")\n" +
			"- ❌ Conflicting instructions in the same prompt\n",
		contentZH: "# Prompt 工程指南\n\n" +
			"编写高效提示词的核心技巧，获得准确、实用的 AI 回复。\n\n" +
			"## 核心技巧\n\n" +
			"### 1. 具体明确\n\n" +
			"| 模糊 ❌ | 具体 ✅ |\n" +
			"|---------|--------|\n" +
			"| \"写一个函数\" | \"写一个 TypeScript 函数，使用正则验证邮箱地址，返回 boolean\" |\n" +
			"| \"改好一点\" | \"改进错误处理：添加 try-catch，返回类型化的错误对象\" |\n\n" +
			"### 2. 提供上下文\n\n" +
			"包含：\n" +
			"- **角色**：\"你是一位资深后端工程师，正在审查 Go 代码\"\n" +
			"- **约束**：\"必须兼容 Node.js 18，不使用外部依赖\"\n" +
			"- **格式**：\"输出为 markdown 表格\"\n" +
			"- **示例**：展示输入/输出示例\n\n" +
			"### 3. 思维链\n\n" +
			"对于复杂问题，要求 AI 逐步思考：\n\n" +
			"\"分析这个数据库架构的性能问题。对每个问题：\n" +
			"1. 识别问题\n" +
			"2. 解释为什么是问题\n" +
			"3. 给出具体的 SQL 修复方案\n" +
			"4. 估算性能影响\"\n\n" +
			"### 4. 少样本示例\n\n" +
			"提供 2-3 个期望输出的示例：\n\n" +
			"\"将需求转换为用户故事：\n\n" +
			"需求：用户可以重置密码\n" +
			"故事：作为用户，我希望通过邮箱重置密码，以便重新访问我的账户。\n\n" +
			"需求：管理员可以封禁用户\n" +
			"故事：作为管理员，我希望通过 ID 封禁用户，以便执行社区规范。\n\n" +
			"现在转换：用户可以将数据导出为 CSV\"\n\n" +
			"### 5. 迭代优化\n" +
			"- 从宽泛的提示开始，逐步缩小范围\n" +
			"- 使用\"保留 X，但修改 Y\"进行增量调整\n" +
			"- 问\"你做了什么假设？\"来发现盲点\n\n" +
			"## 反模式\n\n" +
			"- ❌ 一次给太多指令（分步骤）\n" +
			"- ❌ 模糊的代词（\"让它跟那个配合\"）\n" +
			"- ❌ 没有成功标准（\"做好一点\"）\n" +
			"- ❌ 同一个提示中有互相矛盾的指令\n",
	},
	{
		title: "Dockerfile Best Practices", titleZH: "Dockerfile 最佳实践",
		desc: "Write optimized, secure Dockerfiles for production applications", descZH: "编写优化、安全的生产级 Dockerfile",
		category: "Automation", icon: "📦", tags: "docker,dockerfile,containers,security,optimization", skillType: "general", featured: 0, likes: 4, views: 29,
		content: "# Dockerfile Best Practices\n\n" +
			"Write optimized, secure, and maintainable Dockerfiles.\n\n" +
			"## Multi-Stage Build Template\n\n" +
			cb + "dockerfile\n" +
			"# Stage 1: Build\n" +
			"FROM node:20-alpine AS builder\n" +
			"WORKDIR /app\n" +
			"COPY package*.json ./\n" +
			"RUN npm ci --only=production && \\\n" +
			"    cp -r node_modules prod_modules && \\\n" +
			"    npm ci\n" +
			"COPY . .\n" +
			"RUN npm run build\n\n" +
			"# Stage 2: Production\n" +
			"FROM node:20-alpine\n" +
			"RUN addgroup -g 1001 app && adduser -u 1001 -G app -s /bin/sh -D app\n" +
			"WORKDIR /app\n" +
			"COPY --from=builder /app/prod_modules ./node_modules\n" +
			"COPY --from=builder /app/dist ./dist\n" +
			"USER app\n" +
			"EXPOSE 3000\n" +
			"HEALTHCHECK CMD wget -q --spider http://localhost:3000/health || exit 1\n" +
			"CMD [\"node\", \"dist/server.js\"]\n" +
			cb + "\n\n" +
			"## Optimization Rules\n\n" +
			"| Rule | Why |\n" +
			"|------|-----|\n" +
			"| Use Alpine base images | 5MB vs 900MB |\n" +
			"| Multi-stage builds | Exclude build tools from final image |\n" +
			"| Copy package.json first | Cache dependency layer |\n" +
			"| Combine RUN commands | Fewer layers = smaller image |\n" +
			"| Use `.dockerignore` | Exclude node_modules, .git, tests |\n" +
			"| Pin versions | Reproducible builds |\n\n" +
			"## Security Rules\n\n" +
			"- Run as non-root user (`USER app`)\n" +
			"- Don't store secrets in image (use env vars or secrets manager)\n" +
			"- Scan for vulnerabilities: `docker scout cves`\n" +
			"- Use `COPY` not `ADD` (ADD auto-extracts archives)\n" +
			"- Set `HEALTHCHECK` for container orchestration\n",
		contentZH: "# Dockerfile 最佳实践\n\n" +
			"编写优化、安全、可维护的 Dockerfile。\n\n" +
			"## 多阶段构建模板\n\n" +
			cb + "dockerfile\n" +
			"# 阶段 1：构建\n" +
			"FROM node:20-alpine AS builder\n" +
			"WORKDIR /app\n" +
			"COPY package*.json ./\n" +
			"RUN npm ci --only=production && \\\n" +
			"    cp -r node_modules prod_modules && \\\n" +
			"    npm ci\n" +
			"COPY . .\n" +
			"RUN npm run build\n\n" +
			"# 阶段 2：生产\n" +
			"FROM node:20-alpine\n" +
			"RUN addgroup -g 1001 app && adduser -u 1001 -G app -s /bin/sh -D app\n" +
			"WORKDIR /app\n" +
			"COPY --from=builder /app/prod_modules ./node_modules\n" +
			"COPY --from=builder /app/dist ./dist\n" +
			"USER app\n" +
			"EXPOSE 3000\n" +
			"HEALTHCHECK CMD wget -q --spider http://localhost:3000/health || exit 1\n" +
			"CMD [\"node\", \"dist/server.js\"]\n" +
			cb + "\n\n" +
			"## 优化规则\n\n" +
			"| 规则 | 原因 |\n" +
			"|------|------|\n" +
			"| 使用 Alpine 基础镜像 | 5MB vs 900MB |\n" +
			"| 多阶段构建 | 最终镜像不包含构建工具 |\n" +
			"| 先复制 package.json | 缓存依赖层 |\n" +
			"| 合并 RUN 命令 | 更少层 = 更小镜像 |\n" +
			"| 使用 .dockerignore | 排除 node_modules、.git |\n" +
			"| 固定版本号 | 可重复构建 |\n\n" +
			"## 安全规则\n\n" +
			"- 以非 root 用户运行（`USER app`）\n" +
			"- 不在镜像中存储密钥\n" +
			"- 扫描漏洞：`docker scout cves`\n" +
			"- 使用 `COPY` 而非 `ADD`\n" +
			"- 设置 `HEALTHCHECK`\n",
	},
}
