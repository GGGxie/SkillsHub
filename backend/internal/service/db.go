package service

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbPath string) *sql.DB {
	db, err := sql.Open("sqlite3", dbPath+"?_journal_mode=WAL")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	createTables(db)
	seedData(db)
	return db
}

func createTables(db *sql.DB) {
	queries := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			google_id TEXT UNIQUE,
			email TEXT UNIQUE NOT NULL,
			name TEXT NOT NULL,
			avatar TEXT DEFAULT '',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS skills (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			title_zh TEXT DEFAULT '',
			description TEXT NOT NULL,
			desc_zh TEXT DEFAULT '',
			category TEXT NOT NULL,
			icon TEXT DEFAULT '🤖',
			image TEXT DEFAULT '',
			author_id INTEGER NOT NULL,
			author_name TEXT DEFAULT '',
			author_avatar TEXT DEFAULT '',
			content TEXT DEFAULT '',
			content_zh TEXT DEFAULT '',
			tags TEXT DEFAULT '',
			skill_type TEXT DEFAULT 'general',
			featured INTEGER DEFAULT 0,
			likes INTEGER DEFAULT 0,
			views INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (author_id) REFERENCES users(id)
		)`,
		`CREATE TABLE IF NOT EXISTS articles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			title_zh TEXT DEFAULT '',
			description TEXT NOT NULL,
			desc_zh TEXT DEFAULT '',
			category TEXT DEFAULT 'general',
			content TEXT DEFAULT '',
			content_zh TEXT DEFAULT '',
			author_id INTEGER NOT NULL,
			author_name TEXT DEFAULT '',
			views INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)`,
		`CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			skill_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			user_name TEXT DEFAULT '',
			avatar TEXT DEFAULT '',
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (skill_id) REFERENCES skills(id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		)`,
		`CREATE TABLE IF NOT EXISTS likes (
			user_id INTEGER NOT NULL,
			skill_id INTEGER NOT NULL,
			PRIMARY KEY (user_id, skill_id)
		)`,
	}

	for _, q := range queries {
		if _, err := db.Exec(q); err != nil {
			log.Fatal("Failed to create table:", err)
		}
	}
}

func seedData(db *sql.DB) {
	var count int
	db.QueryRow("SELECT COUNT(*) FROM skills").Scan(&count)
	if count > 0 {
		return
	}

	db.Exec(`INSERT INTO users (id, google_id, email, name, avatar) VALUES (1, 'system', 'system@skillshub.cc', 'SkillsHub', '')`)

	for _, s := range seedSkills {
		db.Exec(`INSERT INTO skills (title, title_zh, description, desc_zh, category, icon, author_id, author_name, content, content_zh, tags, skill_type, featured, likes, views)
			VALUES (?, ?, ?, ?, ?, ?, 1, 'SkillsHub', ?, ?, ?, ?, ?, ?, ?)`,
			s.title, s.titleZH, s.desc, s.descZH, s.category, s.icon, s.content, s.contentZH, s.tags, s.skillType, s.featured, s.likes, s.views)
	}

	articles := []struct {
		title, titleZH, desc, descZH, category, content, contentZH string
		views                                                      int
	}{
		{"Getting Started with AI Skills", "AI技能入门指南",
			"Learn the basics of AI Skills and how to use them effectively", "学习AI技能的基础知识以及如何有效使用它们",
			"Getting Started",
			"# Getting Started with AI Skills\n\n" +
				"AI Skills are modular, reusable instruction sets that enhance AI assistant capabilities. " +
				"Think of them as specialized \"expertise modules\" that you can plug into tools like Cursor, Claude, or ChatGPT.\n\n" +
				"## What is a Skill?\n\n" +
				"A Skill is a structured set of instructions that tells an AI assistant **how** to perform a specific task — " +
				"from code review to data analysis to content writing. Each skill includes:\n\n" +
				"- **Instructions**: Step-by-step guidance\n" +
				"- **Templates**: Output formats and structures\n" +
				"- **Examples**: Concrete input/output pairs\n" +
				"- **Checklists**: Quality verification steps\n\n" +
				"## How to Use Skills\n\n" +
				"### In Cursor IDE\n" +
				"1. Save the skill as a `.md` file in `~/.cursor/skills/` (personal) or `.cursor/skills/` (project)\n" +
				"2. The AI agent automatically discovers and applies relevant skills\n" +
				"3. You can also reference skills manually in your prompts\n\n" +
				"### As Prompts\n" +
				"1. Copy the skill content\n" +
				"2. Paste it as a system prompt or instruction prefix\n" +
				"3. Follow up with your specific task\n\n" +
				"## Tips for Best Results\n\n" +
				"- **Be specific** in your requests — skills work best with clear inputs\n" +
				"- **Iterate** — start with the skill's output and refine\n" +
				"- **Combine skills** — use a code review skill after a code generation skill\n" +
				"- **Customize** — adapt skills to your team's conventions\n",
			"# AI 技能入门指南\n\n" +
				"AI 技能是模块化、可复用的指令集，用于增强 AI 助手的能力。可以把它们想象成专业的\"专长模块\"，" +
				"可以插入 Cursor、Claude 或 ChatGPT 等工具中使用。\n\n" +
				"## 什么是 Skill？\n\n" +
				"Skill 是一组结构化的指令，告诉 AI 助手**如何**执行特定任务 — 从代码审查到数据分析到内容创作。每个 Skill 包含：\n\n" +
				"- **指令**：逐步引导\n" +
				"- **模板**：输出格式和结构\n" +
				"- **示例**：具体的输入/输出对\n" +
				"- **清单**：质量验证步骤\n\n" +
				"## 如何使用\n\n" +
				"### 在 Cursor IDE 中\n" +
				"1. 将 Skill 保存为 `.md` 文件到 `~/.cursor/skills/`（个人）或 `.cursor/skills/`（项目）\n" +
				"2. AI 助手会自动发现并应用相关 Skill\n\n" +
				"### 作为提示词\n" +
				"1. 复制 Skill 内容\n" +
				"2. 粘贴为系统提示或指令前缀\n" +
				"3. 然后输入你的具体任务\n",
			3200},
		{"How to Create Your Own Skill", "如何创建自己的技能",
			"Step-by-step guide to creating and publishing AI Skills", "创建和发布AI技能的分步指南",
			"Getting Started",
			"# How to Create Your Own Skill\n\n" +
				"Creating a skill is straightforward. Follow this guide to build, test, and share your own AI skills.\n\n" +
				"## Step 1: Define the Purpose\n\n" +
				"Answer these questions:\n" +
				"- What specific task does this skill solve?\n" +
				"- Who is the target user? (developers, writers, designers)\n" +
				"- What makes your approach better than a generic prompt?\n\n" +
				"## Step 2: Write the Skill\n\n" +
				"A good skill includes:\n\n" +
				"1. **Clear instructions** — What to do, step by step\n" +
				"2. **Templates** — Structured output formats\n" +
				"3. **Examples** — Show input → output pairs\n" +
				"4. **Checklists** — Verification criteria\n\n" +
				"Keep it concise. Aim for under 500 lines. If you need more detail, link to reference files.\n\n" +
				"## Step 3: Test It\n\n" +
				"Try the skill with different inputs:\n" +
				"- Does it handle edge cases?\n" +
				"- Is the output consistent?\n" +
				"- Would a new user understand how to use it?\n\n" +
				"## Step 4: Share It\n\n" +
				"Share your skill on SkillsHub:\n" +
				"1. Click \"Submit Skill\" on the website\n" +
				"2. Fill in the title, description, and content\n" +
				"3. Tag it with relevant categories\n" +
				"4. Publish and get feedback from the community!\n",
			"# 如何创建自己的技能\n\n" +
				"创建 Skill 很简单。按照本指南构建、测试和分享你自己的 AI 技能。\n\n" +
				"## 步骤 1：定义目的\n\n" +
				"回答这些问题：\n" +
				"- 这个 Skill 解决什么具体任务？\n" +
				"- 目标用户是谁？\n" +
				"- 你的方法比通用提示词好在哪里？\n\n" +
				"## 步骤 2：编写 Skill\n\n" +
				"好的 Skill 包含：清晰的指令、模板、示例和检查清单。\n\n" +
				"## 步骤 3：测试\n\n" +
				"用不同的输入测试，确保处理边界情况，输出一致。\n\n" +
				"## 步骤 4：分享\n\n" +
				"在 SkillsHub 上分享你的 Skill，获得社区反馈！\n",
			2800},
		{"Claude Skills vs MCP", "Claude技能 vs MCP",
			"Understanding the differences between Claude Skills and MCP", "了解Claude技能和MCP之间的区别",
			"Integration",
			"# Claude Skills vs MCP\n\n" +
				"Both Skills and MCP (Model Context Protocol) extend AI capabilities, but they work differently.\n\n" +
				"## Quick Comparison\n\n" +
				"| Aspect | Skills | MCP |\n" +
				"|--------|--------|-----|\n" +
				"| What | Instruction sets (text) | Tool/API connections (code) |\n" +
				"| How | Prompt-based guidance | Function calling interface |\n" +
				"| Setup | Drop a .md file | Run a server process |\n" +
				"| Capabilities | Knowledge, templates, workflows | External actions (DB, API, file system) |\n" +
				"| Complexity | Low (just Markdown) | Medium (requires server) |\n\n" +
				"## When to Use Skills\n\n" +
				"- Teaching the AI **how** to approach a task\n" +
				"- Encoding team conventions and best practices\n" +
				"- Providing templates and output formats\n" +
				"- Knowledge that applies across many tasks\n\n" +
				"## When to Use MCP\n\n" +
				"- The AI needs to **perform actions** (query database, call API)\n" +
				"- Real-time data access is required\n" +
				"- Integrating with external services\n" +
				"- Complex tool chains with state\n\n" +
				"## Best of Both Worlds\n\n" +
				"Combine them! Use a Skill to define the workflow and an MCP server to execute it:\n\n" +
				"- **Skill**: \"When reviewing a PR, check for security issues, run tests, and verify coverage\"\n" +
				"- **MCP**: GitHub API to fetch PR diffs, test runner to execute tests, coverage tool to check metrics\n",
			"# Claude 技能 vs MCP\n\n" +
				"Skills 和 MCP（模型上下文协议）都能扩展 AI 能力，但工作方式不同。\n\n" +
				"## 快速对比\n\n" +
				"| 方面 | Skills | MCP |\n" +
				"|------|--------|-----|\n" +
				"| 是什么 | 指令集（文本） | 工具/API 连接（代码） |\n" +
				"| 如何工作 | 基于提示的引导 | 函数调用接口 |\n" +
				"| 设置 | 放置 .md 文件 | 运行服务器进程 |\n" +
				"| 能力 | 知识、模板、工作流 | 外部操作（数据库、API） |\n\n" +
				"## 何时使用 Skills\n\n" +
				"- 教 AI **如何**处理任务\n" +
				"- 编码团队规范和最佳实践\n" +
				"- 提供模板和输出格式\n\n" +
				"## 何时使用 MCP\n\n" +
				"- AI 需要**执行操作**（查询数据库、调用 API）\n" +
				"- 需要实时数据访问\n" +
				"- 集成外部服务\n\n" +
				"## 两全其美\n\n" +
				"结合使用！用 Skill 定义工作流，用 MCP 服务器执行它。\n",
			1950},
	}

	for _, a := range articles {
		db.Exec(`INSERT INTO articles (title, title_zh, description, desc_zh, category, content, content_zh, author_id, author_name, views)
			VALUES (?, ?, ?, ?, ?, ?, ?, 1, 'SkillsHub', ?)`,
			a.title, a.titleZH, a.desc, a.descZH, a.category, a.content, a.contentZH, a.views)
	}
}
