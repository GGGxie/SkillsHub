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

	skills := []struct {
		title, titleZH, desc, descZH, category, icon, tags, skillType string
		featured                                                      int
		likes, views                                                  int
	}{
		{"Code Review Assistant", "代码审查助手", "Automated code review with best practices and security checks", "自动化代码审查，包含最佳实践和安全检查", "Coding & Development", "🔍", "code,review,security", "ai-powered", 1, 128, 2340},
		{"Blog Post Writer", "博客文章写手", "Generate well-structured blog posts with SEO optimization", "生成结构良好的博客文章，包含SEO优化", "Writing & Content", "✍️", "writing,blog,seo", "general", 1, 96, 1850},
		{"Data Analysis Pipeline", "数据分析流水线", "Automated data analysis with visualization and reporting", "自动化数据分析，包含可视化和报告", "Data & Analytics", "📊", "data,analysis,visualization", "ai-powered", 1, 84, 1620},
		{"API Integration Helper", "API集成助手", "Simplify REST API integration with auto-generated client code", "简化REST API集成，自动生成客户端代码", "Coding & Development", "🔗", "api,integration,rest", "general", 0, 72, 1430},
		{"Email Template Designer", "邮件模板设计师", "Create responsive email templates with dynamic content", "创建响应式邮件模板，支持动态内容", "Design & Creative", "📧", "email,template,design", "general", 0, 65, 1280},
		{"Test Case Generator", "测试用例生成器", "Auto-generate comprehensive test cases from specifications", "根据规格说明自动生成全面的测试用例", "Coding & Development", "🧪", "testing,automation,quality", "ai-powered", 1, 110, 2100},
		{"Translation Assistant", "翻译助手", "Multi-language translation with context awareness", "多语言翻译，支持上下文感知", "Writing & Content", "🌐", "translation,language,i18n", "ai-powered", 0, 58, 980},
		{"Task Automation Bot", "任务自动化机器人", "Automate repetitive tasks with customizable workflows", "自动化重复任务，支持可自定义工作流", "Automation", "🤖", "automation,workflow,bot", "general", 0, 92, 1760},
		{"UI Component Builder", "UI组件构建器", "Generate React/Vue components from design descriptions", "根据设计描述生成React/Vue组件", "Design & Creative", "🎨", "ui,components,react,vue", "ai-powered", 1, 145, 2680},
		{"Database Schema Designer", "数据库架构设计师", "Design optimal database schemas with relationship mapping", "设计最优数据库架构，支持关系映射", "Coding & Development", "🗄️", "database,schema,design", "general", 0, 78, 1350},
		{"SEO Optimizer", "SEO优化器", "Analyze and optimize content for search engines", "分析和优化搜索引擎内容", "Writing & Content", "🔎", "seo,optimization,content", "ai-powered", 0, 67, 1150},
		{"DevOps Pipeline Setup", "DevOps流水线配置", "Configure CI/CD pipelines with best practices", "配置CI/CD流水线，包含最佳实践", "Automation", "⚙️", "devops,cicd,pipeline", "general", 0, 88, 1580},
	}

	for _, s := range skills {
		db.Exec(`INSERT INTO skills (title, title_zh, description, desc_zh, category, icon, author_id, author_name, tags, skill_type, featured, likes, views)
			VALUES (?, ?, ?, ?, ?, ?, 1, 'SkillsHub', ?, ?, ?, ?, ?)`,
			s.title, s.titleZH, s.desc, s.descZH, s.category, s.icon, s.tags, s.skillType, s.featured, s.likes, s.views)
	}

	articles := []struct {
		title, titleZH, desc, descZH, category, content, contentZH string
		views                                                      int
	}{
		{"Getting Started with AI Skills", "AI技能入门指南", "Learn the basics of AI Skills and how to use them effectively", "学习AI技能的基础知识以及如何有效使用它们", "Getting Started",
			"# Getting Started\n\nAI Skills are modular capabilities that enhance AI functionality...", "# 入门指南\n\nAI技能是增强AI功能的模块化能力...", 3200},
		{"How to Create Your Own Skill", "如何创建自己的技能", "Step-by-step guide to creating and publishing AI Skills", "创建和发布AI技能的分步指南", "Getting Started",
			"# Creating Skills\n\nFollow these steps to create your own skill...", "# 创建技能\n\n按照以下步骤创建你自己的技能...", 2800},
		{"Claude Skills vs MCP", "Claude技能 vs MCP", "Understanding the differences between Claude Skills and MCP", "了解Claude技能和MCP之间的区别", "Integration",
			"# Claude Skills vs MCP\n\nBoth are powerful tools for AI augmentation...", "# Claude技能 vs MCP\n\n两者都是强大的AI增强工具...", 1950},
	}

	for _, a := range articles {
		db.Exec(`INSERT INTO articles (title, title_zh, description, desc_zh, category, content, content_zh, author_id, author_name, views)
			VALUES (?, ?, ?, ?, ?, ?, ?, 1, 'SkillsHub', ?)`,
			a.title, a.titleZH, a.desc, a.descZH, a.category, a.content, a.contentZH, a.views)
	}
}
