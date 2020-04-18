#include <string>
#include <vector>
#include <iostream>
#include <cstdio>
#include <fstream>
#include <cstring>

using namespace std;

class mycsv {
private:
    FILE* fp;
    int bufsize;
    char* line;
    char comma;
    void init() {
        fp = NULL;
        line = new char[bufsize];
        comma = ',';
    }
public:
    // set comma charcter
    void setcomma(char c) {
        comma = c;
    }
    // open csv file
    bool open(string filename) {
        if (fp != NULL) fclose(fp);
        fp = fopen(filename.c_str(), "r");
        if (fp == NULL) return false;
        return true;
    }
    // close csv file
    void close() {
        if (fp != NULL) fclose(fp);
        fp = NULL;
    }
    vector<string> loadnextline() {
        memset(line, bufsize, sizeof(char) * bufsize);
        fgets(line, bufsize, fp);

        size_t len = strlen(line);
        line[len - 1] = comma;
        line[len] = '\0';

        vector<string> cells;
        string l;

        for (int i = 0; i < len; i++) {
            if (line[i] == '\r') continue;
            if (line[i] == comma) {
                cells.push_back(l);
                l.clear();
                continue;
            }
            l.push_back(line[i]);
        }
        return cells;
    }
    bool eof() {
        if (fp == NULL) return true;
        return feof(fp);
    }
    vector<vector<string>> loadall() {
        vector<vector<string>> ans;
        while (!eof())
        {
            ans.push_back(loadnextline());
        }
        return ans;
    }
    mycsv(string filename, int linemax = 4096) {
        bufsize = linemax;
        init();
        open(filename);
    }
    mycsv(int linemax) {
        bufsize = linemax;
        init();
    }
    ~mycsv() {
        delete[] line;
        if (fp == NULL) return;
        fclose(fp);
        fp = NULL;
    }
};
#include <unordered_map>
#include <unordered_set>

typedef pair<string, string> node;

int main() {

    unordered_set <string> g_street_wayids;

    // ノード情報の読み取り
    mycsv csv("node.csv");
    unordered_map<string, node> nodes;
    while (!csv.eof())
    {
        vector<string> cells = csv.loadnextline();
        if (cells.size() < 3)continue;
        nodes[cells[0]] = node(cells[1], cells[2]);
    }

    // タグ情報の読み取り
    csv.open("tag.csv");
    while (!csv.eof())
    {
        vector<string> cells = csv.loadnextline();
        if (cells.size() < 3)continue;
        if (cells[1] == "highway") g_street_wayids.insert(cells[0]);
        else if (cells[1] == "highway") g_street_wayids.insert(cells[0]);
    }

    // way情報の読み取り
    ofstream ofs("combind.csv");
    csv.open("edge.csv");
    while (!csv.eof())
    {
        vector<string> cells = csv.loadnextline();
        if (cells.size() < 3)continue;
        if (nodes.find(cells[1]) == nodes.end()) continue;
        if (nodes.find(cells[2]) == nodes.end()) continue;
        if (g_street_wayids.find(cells[0]) == g_street_wayids.end()) continue;
        ofs << cells[0] << ',' << cells[1] << ',' << nodes[cells[1]].first << ',' << nodes[cells[1]].second << ',' << cells[2] << ',' << nodes[cells[2]].first << ',' << nodes[cells[2]].second << endl;
    }
}