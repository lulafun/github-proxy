document.addEventListener('DOMContentLoaded', function() {
    const searchForm = document.getElementById('search-form');
    const searchInput = document.getElementById('search-input');

    // 处理表单提交
    searchForm.addEventListener('submit', function(e) {
        e.preventDefault();

        const searchValue = searchInput.value.trim();
        if (!searchValue) {
            // 空输入时添加震动效果
            searchInput.classList.add('shake');
            setTimeout(() => {
                searchInput.classList.remove('shake');
            }, 500);
            return;
        }

        // 清理URL格式
        let url = searchValue;
        if (!url.startsWith('http://') && !url.startsWith('https://')) {
            url = 'https://' + url;
        }

        // 检查是否为GitHub URL
        const githubPattern = /^https?:\/\/(github\.com|raw\.githubusercontent\.com|gist\.github\.com)/i;
        if (!githubPattern.test(url)) {
            alert('Please enter a valid GitHub URL');
            return;
        }

        // 重定向到处理路径
        window.location.href = '/' + url;
    });

    // 添加键盘快捷键支持
    document.addEventListener('keydown', function(e) {
        // 按下/键时聚焦到搜索框
        if (e.key === '/' && document.activeElement !== searchInput) {
            e.preventDefault();
            searchInput.focus();
        }
    });
});

// 添加CSS动画
const style = document.createElement('style');
style.textContent = `
    @keyframes shake {
        0%, 100% { transform: translateX(0); }
        10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
        20%, 40%, 60%, 80% { transform: translateX(5px); }
    }
    
    .shake {
        animation: shake 0.5s;
    }
`;
document.head.appendChild(style);